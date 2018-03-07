package context

import (
	"net"
	"testing"

	"github.com/diltram/go-telnet/telsh"
	"github.com/sirupsen/logrus"
)

func TestNewContext(t *testing.T) {
	ctx := NewContext()

	if ctx.Logger() != nil {
		t.Error("Context doesn't suppose to have a logger configured")
	}

	if ctx.Conn() != nil {
		t.Error("Context doesn't suppose to have a connection configured")
	}

	if ctx.Handler() != nil {
		t.Error("Context doesn't suppose to have a base handler configured")
	}
}

func TestInjectHandler(t *testing.T) {
	ctx := NewContext()
	handler := telsh.NewShellHandler()

	ctx.InjectHandler(handler)

	if ctx.Handler() != handler {
		t.Error("Received different handler that injected")
	}
}

func TestInjectLogger(t *testing.T) {
	ctx := NewContext()
	logger := logrus.New()

	ctx.InjectLogger(logger)

	if ctx.Logger() != logger {
		t.Error("Received different logger that injected")
	}
}

func TestInjectConn(t *testing.T) {
	ctx := NewContext()
	conn, _ := net.Dial("tcp", ":80")

	ctx.InjectConn(conn)

	if ctx.Conn() != conn {
		t.Error("Received different connection that injected")
	}
}
