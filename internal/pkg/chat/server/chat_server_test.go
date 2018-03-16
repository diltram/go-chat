//@TODO: Add way better tests
package server

import (
	"testing"

	"github.com/diltram/go-chat/internal/pkg/server/context"
)

func TestNewServer(t *testing.T) {
	addr := ":telnet"
	ctx := new(context.Context)
	srv := NewServer(addr, ctx).(*ChatServer)

	if srv.Addr != addr {
		t.Errorf("Different address configured, expected %s, actual %s", addr, srv.Addr)
	}
}
