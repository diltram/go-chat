//@TODO: Add way better tests
package server

import (
	"testing"

	"github.com/diltram/go-chat/internal/pkg/server/context"
	"github.com/diltram/go-chat/internal/pkg/server/handler"
)

func TestNewServer(t *testing.T) {
	addr := ":telnet"
	var ctx context.Context
	var handler handler.Handler
	srv := NewServer(addr, handler, ctx).(*ChatServer)

	if srv.Addr != addr {
		t.Errorf("Different address configured, expected %s, actual %s", addr, srv.Addr)
	}
}
