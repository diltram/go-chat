package context

import (
	"net"

	"github.com/diltram/go-telnet"
)

// HandlerContext embedds base Context from go-telnet library and extends it
// with support for handler. It allows to inject and retrieve base handler
// which is used for handling requests from clients.
type HandlerContext interface {
	telnet.Context

	Handler() telnet.Handler
	InjectHandler(telnet.Handler) telnet.Context
}

type chatContext struct {
	logger  telnet.Logger
	conn    net.Conn
	handler telnet.Handler
}

// NewContext initializes base context for commands handler.
func NewContext() HandlerContext {
	ctx := chatContext{}

	return &ctx
}

// Logger provides access to preconfigured logging mechanism.
func (ctx *chatContext) Logger() telnet.Logger {
	return ctx.logger
}

// InjectLogger let's configure in user context specific logging mechanism.
func (ctx *chatContext) InjectLogger(logger telnet.Logger) telnet.Context {
	ctx.logger = logger

	return ctx
}

// Users connection via which he's sending messages.
func (ctx *chatContext) Conn() net.Conn {
	return ctx.conn
}

// InjectConn provides functionality to attach user's connection data into
// context.
func (ctx *chatContext) InjectConn(conn net.Conn) telnet.Context {
	ctx.conn = conn

	return ctx
}

// InjectHandler allows to provide base handler used for serving all the
// requests. It let's to interact with handler in commands e.g. send message to
// all users connected to server.
func (ctx *chatContext) InjectHandler(handler telnet.Handler) telnet.Context {
	ctx.handler = handler

	return ctx
}

// Handler returns base handler with access to all it's data and methods.
func (ctx *chatContext) Handler() telnet.Handler {
	return ctx.handler
}
