package command

import (
	"bytes"
	"strings"
	"testing"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/chat/context"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

func TestCallNick(t *testing.T) {
	cmd := NickCommand{}
	ctx := usrctx.NewUserContext()

	chat := chat.NewChat()
	ctx.SetAttribute("chat", chat)

	usr, server, closer := user.MockUser()
	defer closer()

	readBuf := new(bytes.Buffer)
	go func() {
		buf := make([]byte, 20)
		for {
			server.Read(buf)
			readBuf.Write(buf)
		}
	}()
	ctx.SetAttribute("user", usr)

	channel1 := channel.NewChannel("channel1")
	channel1.AddUser(usr)
	chat.AddChannel(channel1)
	ctx.SetAttribute("channel", channel1)

	cases := []struct {
		Buffer string
		Nick   string
	}{
		{"/nick nick", "nick"},
		{"/nick i9)(*()#$", "i9)(*()#$"},
	}

	for _, tt := range cases {
		buf := bytes.NewBufferString(tt.Buffer)
		fields := strings.Fields(tt.Buffer)
		cmd.Call(ctx, fields, buf)

		if readBuf.Len() != 0 {
			t.Errorf("There shouldn't be any data writen to user. Data: %s", readBuf.String())
		}

		if usr.Name() != tt.Nick {
			t.Errorf("User nick: expecting %s, actual %s", tt.Nick, usr.Name())
		}
	}
}

func TestCallNickNoData(t *testing.T) {
	cmd := NickCommand{}
	ctx := usrctx.NewUserContext()

	chat := chat.NewChat()
	ctx.SetAttribute("chat", chat)

	usr, server, closer := user.MockUser()
	defer closer()

	readBuf := new(bytes.Buffer)
	go func() {
		buf := make([]byte, 20)
		for {
			server.Read(buf)
			readBuf.Write(buf)
		}
	}()
	ctx.SetAttribute("user", usr)

	channel1 := channel.NewChannel("channel1")
	channel1.AddUser(usr)
	chat.AddChannel(channel1)
	ctx.SetAttribute("channel", channel1)

	cases := []struct {
		Buffer string
	}{
		{"/nick"},
		{"/nick "},
	}

	for _, tt := range cases {
		buf := bytes.NewBufferString(tt.Buffer)
		fields := strings.Fields(tt.Buffer)
		oldNick := usr.Name()
		cmd.Call(ctx, fields, buf)

		if !bytes.Contains(readBuf.Bytes(), []byte("Cannot change a nick to empty string, nice try...")) {
			t.Error("Missing information that nick cannot be changed")
		}

		if usr.Name() != oldNick {
			t.Errorf("User nick: expecting %s, actual %s", oldNick, usr.Name())
		}
		readBuf.Reset()
	}
}
