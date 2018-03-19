package usrctx

import (
	"testing"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

func TestNoChat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx := NewUserContext()
	ctx.Chat()
}

func TestWrongChat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx := NewUserContext()
	ctx.SetAttribute("chat", "value")
	ctx.Chat()
}

func TestChat(t *testing.T) {
	ctx := NewUserContext()
	chatInst := chat.NewChat()
	ctx.SetAttribute("chat", chatInst)
	actual := ctx.Chat()

	if actual != chatInst {
		t.Errorf("Get chat, expected %v, actual %v", chatInst, actual)
	}
}

func TestNoUser(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx := NewUserContext()
	ctx.User()
}

func TestWrongUser(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx := NewUserContext()
	ctx.SetAttribute("user", "value")
	ctx.User()
}

func TestUser(t *testing.T) {
	ctx := NewUserContext()
	usr, _, closer := user.MockUser()
	defer closer()

	ctx.SetAttribute("user", usr)
	actual := ctx.User()

	if actual != usr {
		t.Errorf("Get user, expected %v, actual %v", usr, actual)
	}
}

func TestNoChannel(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx := NewUserContext()
	ctx.Channel()
}

func TestWrongChannel(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx := NewUserContext()
	ctx.SetAttribute("channel", "value")
	ctx.Channel()
}

func TestChannel(t *testing.T) {
	ctx := NewUserContext()
	ch := channel.NewChannel("test")
	ctx.SetAttribute("channel", ch)
	actual := ctx.Channel()

	if actual != ch {
		t.Errorf("Get channel, expected %v, actual %v", ch, actual)
	}
}
