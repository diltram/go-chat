package server

import (
	"fmt"
	"io"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/user"
	"github.com/diltram/go-chat/internal/pkg/server"
	"github.com/diltram/go-chat/internal/pkg/server/context"
	"github.com/diltram/go-chat/internal/pkg/server/handler"
)

// ListenAndServe allows in one command to create a server and start it on
// specified address.
func ListenAndServe(addr string) error {
	srv := NewServer(addr, context.NewContext())
	return srv.ListenAndServe()
}

// NewServer creates new ChatServer with specified address to listen on.
func NewServer(addr string, ctx context.Context) server.Server {
	return &ChatServer{
		Addr: addr,
		ctx:  ctx,
	}
}

// NewServerHandler creates new ChatServer with specified address to listen on.
func NewServerHandler(addr string, h handler.Handler, ctx context.Context) server.Server {
	return &ChatServer{
		Addr:    addr,
		ctx:     ctx,
		handler: h,
	}
}

// ChatServer implements Server and provides base chat server functionality.
//
// It listens on listener l for new connections which then registers as new
// user in Chat structure.
type ChatServer struct {
	// Addr represents TCP address to listen on. When empty ":telnet" will be
	// used.
	Addr string
	// Handler which will be used for execution of user requests.
	handler handler.Handler
	// Ctx represent internal server context. You can configure anything in
	// here. It will be provided to handler to allow access to data in it.
	ctx context.Context
}

// ListenAndServe creates new listener on which it's starting listening.
func (cs *ChatServer) ListenAndServe() error {
	addr := cs.Addr
	if "" == addr {
		addr = ":telnet"
	}

	listener, err := net.Listen("tcp", addr)
	if nil != err {
		return err
	}

	return cs.Serve(listener)
}

// Serve starts listening on l and send all new connections into Handle method.
func (cs *ChatServer) Serve(l net.Listener) error {
	defer l.Close()

	var tempDelay time.Duration // how long to sleep on accept failure

	for {
		// Wait for a new TELNET client connection.
		log.Debugf("Listening at %q.", l.Addr())
		conn, err := l.Accept()
		if err != nil {
			// That code is taken from Golang HTTP package from Serve method.
			if _, ok := err.(net.Error); ok {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Info("Server: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
		}
		log.Debugf("Received new connection from %q.", conn.RemoteAddr())

		// Handle the new TELNET client connection by spawning a new goroutine.
		go cs.Handle(cs.ctx, conn)
		log.Debugf("Spawned handler to handle connection from %q.", conn.RemoteAddr())
	}
}

// Handle pre-configures all required data and sends request in new goroutine
// to handler for processing.
//
// When context is provided it will be used instead of internal server context.
func (cs *ChatServer) Handle(ctx context.Context, c net.Conn) {
	defer c.Close()

	if ctx == nil {
		ctx = cs.ctx
	}

	if cs.handler == nil {
		log.Debug("Defaulted handler to EchoHandler.")
		cs.handler = handler.EchoHandler{}
	}

	attr, err := ctx.Attribute("chat")
	if err != nil {
		panic(err)
	}

	chatInst, ok := attr.(*chat.Chat)
	if ok == false {
		panic("Chat instance is not available")
	}

	defaultChannel := chatInst.Channels()["default"]
	usr := user.NewUser(c, "")
	defaultChannel.AddUser(usr)
	usrCtx := cs.ctx.Clone()
	usrCtx.SetAttribute("user", usr)

	writer := c.(io.Writer)
	reader := c.(io.Reader)

	cs.handler.Serve(usrCtx, writer, reader)

	// User disconnected. Send info about that
	log.Debugf("Closing connection from %q.", c.RemoteAddr())
	chatInst.Channels()["default"].DelUser(usr)
	msg := defaultChannel.AddNotification(usr, fmt.Sprintf("User %s disconnected from channel %s\r\n", usr.Name(), defaultChannel.Name()))
	defaultChannel.SendMessage(usr, msg)
}

// Shutdown gracefully shut downs the server.
//
// For now just fall back to Close and stop server.
// When context is provided it will be used instead of internal server context.
func (cs *ChatServer) Shutdown(ctx context.Context) error {
	//@TODO: Add real code for graceful shutdown
	return cs.Close()
}

// Close stops the server without any checks.
//
// It closes the listener and disconnects all established sessions.
func (cs *ChatServer) Close() error {
	//@TODO: Add real close code
	return nil
}
