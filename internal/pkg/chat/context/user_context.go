package usrctx

import (
	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/server/context"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

// NewUserContext creates a context with access to special fields based on the
// attributes.
func NewUserContext() *UserContext {
	return &UserContext{
		context.NewContext(),
	}
}

// UserContext is a special context wrapper.
// That context provides direct access to Chat, User and Channel attributes.
// It panics whenever any of them is no available.
type UserContext struct {
	context.Context
}

// Chat provides access to attribute.
// It converts default interface{} return into a Chat struct.
// It panics when attribute is not available or when it has a wrong type.
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

// User provides access to attribute.
// It converts default interface{} return into a User struct.
// It panics when attribute is not available or when it has a wrong type.
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

// Channel provides access to attribute.
// It converts default interface{} return into a Channel struct.
// It panics when attribute is not available or when it has a wrong type.
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
