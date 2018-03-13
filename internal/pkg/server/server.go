package server

import (
	"net"

	"github.com/diltram/go-chat/internal/pkg/server/context"
)

// Server interface defines all methods which server needs to support.
type Server interface {
	// ListenAndServe listens on the TCP network address and then calls Serve
	// to handle requests on incoming connections. ListenAndServe always
	// returns a non-nil error.
	ListenAndServe() error
	// Serve accepts incoming connections on the Listener l, creating a new
	// service goroutine for each. The service goroutines read requests and
	// then call handler to reply to them.
	Serve(l net.Listener) error
	//handle(c net.Conn, handler Handler)
	//Shutdown gracefully shuts down the server.
	Shutdown(ctx context.Context) error
	// Stop server without graceful awaiting for any operations to complete.
	Close() error
}
