package server

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"testing"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/server/handler"
	"github.com/diltram/go-chat/internal/pkg/server/context"
)

func TestNewServer(t *testing.T) {
	addr := ":telnet"
	ctx := context.NewContext()
	handler := handler.NewChatHandler()
	srv := NewServer(addr, handler, ctx).(*ChatServer)

	if srv.Addr != addr {
		t.Errorf("Different address configured, expected %s, actual %s", addr, srv.Addr)
	}
}

func TestServer(t *testing.T) {
	ctx := context.NewContext()
	chatInst := chat.NewChat()
	ctx.SetAttribute("chat", chatInst)
	handler := handler.NewChatHandler()

	addr := "127.0.0.1:5555"
	srv := NewServer(addr, handler, ctx)
	go srv.ListenAndServe()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Errorf("Failed to establish connection with server, err %s", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	actual, _, err := reader.ReadLine()
	actual, _, err = reader.ReadLine()
	actual, _, err = reader.ReadLine()
	if err != nil {
		t.Errorf("Failed to read data from conn, err %s", err)
	}

	expected := ", welcome in a channel default. There is 0 other users"
	if !bytes.Contains(actual, []byte(expected)) {
		t.Errorf("Wrong welcome message. Actual %s", actual)
	}

	// Disconnect from server
	io.WriteString(conn, "/quit\r\n")
	actual, _, err = reader.ReadLine()
	actual, _, err = reader.ReadLine()
	if err != nil {
		t.Errorf("Failed to read data from conn, err %s", err)
	}

	expected = "Goodbye!"
	if !bytes.Contains(actual, []byte(expected)) {
		t.Errorf("Wrong goodbye message. Actual %s", actual)
	}

	actualCount := len(chatInst.Channels()["default"].Users())
	if actualCount != 0 {
		t.Errorf("User left channel. Should be dropped from list, count %d", actualCount)
	}
}
