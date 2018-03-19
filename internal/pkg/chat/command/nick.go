package command

import (
	"bytes"
	"fmt"
	"io"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// NickCommand allows to change user's nick.
type NickCommand struct{}

// Name returns a descriptive name of the command used for help.
func (cmd NickCommand) Name() string {
	return "Change nick"
}

// Desc will be displayed in help as complete description of command.
func (cmd NickCommand) Desc() string {
	return "Commands allow to change current nick into a new one."
}

// Cmds return a slice of names which can be used to call that method.
// All of them will be mapped and user can use any of them to trigger
// operation.
func (cmd NickCommand) Cmds() []string {
	return []string{"/n", "/nick"}
}

// Call changes user's nick.
// When nick is not provided it will show error. In opposite situation it will
// change a nick and inform all the users on the channel about that.
func (cmd NickCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	usr := ctx.User()
	chann := ctx.Channel()

	nick := RemoveCmd(fields[0], cmdLine.String())
	if len(nick) == 0 {
		io.WriteString(usr, "Cannot change a nick to empty string, nice try...\r\n")
		return
	}

	oldNick := usr.Name()
	usr.SetName(nick)

	msg := chann.AddNotification(usr, fmt.Sprintf("User %s changed his nick to %s", oldNick, usr.Name()))
	chann.SendMessage(usr, msg)
}

func init() {
	GetRegistry().RegisterCmd(NickCommand{})
}
