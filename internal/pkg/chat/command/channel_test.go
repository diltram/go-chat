package command

import (
	"bytes"
	"testing"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/chat/context"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

func TestCall(t *testing.T) {
	cmd := ChannelCommand{}
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
	channel2 := channel.NewChannel("channel2")
	chat.AddChannel(channel1)
	chat.AddChannel(channel2)
	channel1.AddUser(usr)
	ctx.SetAttribute("channel", channel1)

	cases := []struct {
		Buffer  string
		Fields  []string
		Channel string
		Count   int
	}{
		{"/channel channel2", []string{"/channel", "channel2"}, "channel2", 3},
		{"/channel channel3", []string{"/channel", "channel3"}, "channel3", 4},
	}

	for _, tt := range cases {
		buf := bytes.NewBufferString(tt.Buffer)
		fields := tt.Fields
		cmd.Call(ctx, fields, buf)

		if len(ctx.Chat().Channels()) != tt.Count {
			t.Errorf("Channels list: expecting %d, actual %d", tt.Count, len(ctx.Chat().Channels()))
		}

		if ctx.Channel().Name() != tt.Channel {
			t.Errorf("Channel should be changed to %s, actual %s", tt.Channel, ctx.Channel().Name())
		}
	}
}

func TestCallNoChannel(t *testing.T) {
	cmd := ChannelCommand{}
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
	chat.AddChannel(channel1)
	channel1.AddUser(usr)
	ctx.SetAttribute("channel", channel1)

	cases := []struct {
		Buffer string
		Fields []string
	}{
		{"/channel", []string{"/channel"}},
		{"/channel ", []string{"/channel"}},
	}

	for _, tt := range cases {
		buf := bytes.NewBufferString(tt.Buffer)
		fields := tt.Fields
		cmd.Call(ctx, fields, buf)

		if !bytes.Contains(readBuf.Bytes(), []byte("You need to specify a channel...")) {
			t.Error("Channel should be required, something went wrong")
		}

		if ctx.Channel() != channel1 {
			t.Error("Channel shouldn't be changed")
		}
	}
}
