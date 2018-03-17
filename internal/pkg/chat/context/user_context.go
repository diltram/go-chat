package usrctx

import (
	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/server/context"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

type UserContext struct {
	context.Context
}

func (ctx UserContext) Chat() *chat.Chat {
	attr, err := ctx.Attribute("chat")
	if err != nil {
		panic(err)
	}

	chatInst, ok := attr.(*chat.Chat)
	if ok == false {
		panic("Chat instance is not available")
	}

	return chatInst
}

func (ctx *UserContext) User() *user.User {
	attr, err := ctx.Attribute("user")

	if err != nil {
		panic(err)
	}

	usr, ok := attr.(*user.User)
	if ok == false {
		panic("User instance is not available")
	}

	return usr
}

func (ctx *UserContext) Channel() *channel.Channel {
	attr, err := ctx.Attribute("channel")

	if err != nil {
		panic(err)
	}

	ch, ok := attr.(*channel.Channel)
	if ok == false {
		panic("User's channel is not available")
	}

	return ch
}
