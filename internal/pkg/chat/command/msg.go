package command

import (
	"bytes"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// MsgCommand is a default command.
//
// When none of they keys are found we're assuming that user is sending a
// message and that command will be triggered. Because of that it doesn't
// provide any additional data. It will never appear in help command.
type MsgCommand struct{}

// Name returns just empty string.
func (cmd MsgCommand) Name() string {
	return ""
}

// Desc returns just empty string.
func (cmd MsgCommand) Desc() string {
	return ""
}

// Cmds returns nothing as it will be registered as default command.
func (cmd MsgCommand) Cmds() []string {
	return nil
}

// Call sends message to all the users on specific channel except sender.
func (cmd MsgCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	ch := ctx.Channel()
	usr := ctx.User()

	msg := ch.AddMessage(usr, cmdLine.String())
	ch.SendMessage(usr, msg)
}

func init() {
	GetRegistry().RegisterDefCmd(MsgCommand{})
}
