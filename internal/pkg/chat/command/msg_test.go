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

func TestCallMsg(t *testing.T) {
	cmd := MsgCommand{}
	ctx := usrctx.NewUserContext()

	chat := chat.NewChat()
	ctx.SetAttribute("chat", chat)

	usr, _, closer := user.MockUser()
	defer closer()
	usr2, server2, closer2 := user.MockUser()
	defer closer2()

	readBuf := new(bytes.Buffer)
	go func() {
		buf := make([]byte, 20)
		for {
			server2.Read(buf)
			readBuf.Write(buf)
		}
	}()
	ctx.SetAttribute("user", usr)

	channel1 := channel.NewChannel("channel1")
	channel1.AddUser(usr)
	channel1.AddUser(usr2)
	chat.AddChannel(channel1)
	ctx.SetAttribute("channel", channel1)

	cases := []struct {
		Buffer string
	}{
		{"some text"},
		{"0938q 09q8)%*#()*$# $#@)(*$#)(l"},
	}

	for _, tt := range cases {
		buf := bytes.NewBufferString(tt.Buffer)
		fields := strings.Fields(tt.Buffer)
		cmd.Call(ctx, fields, buf)

		if !bytes.Contains(readBuf.Bytes(), []byte(tt.Buffer)) {
			t.Errorf("Msg: expected %s, actual %s", tt.Buffer, readBuf.String())
		}
		readBuf.Reset()
	}
}
