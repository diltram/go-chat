package command

import (
	"bytes"
	"fmt"
	"io"

	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// ChannelCommand allows to change a channel.
type ChannelCommand struct{}

// Name returns a descriptive name of the command used for help.
func (cmd ChannelCommand) Name() string {
	return "Change channel"
}

// Desc will be displayed in help as complete description of command.
func (cmd ChannelCommand) Desc() string {
	return "Commands allow to change current channel into a new one."
}

// Cmds return a slice of names which can be used to call that method.
// All of them will be mapped and user can use any of them to trigger
// operation.
func (cmd ChannelCommand) Cmds() []string {
	return []string{"/channel", "/chan", "/ch"}
}

// Call executes a process of change of channel.
// When an empty command specified it will show error that channel can't be
// changed.
// It will trigger additional messages on both channels informing people that
// user left/joined a channel.
func (cmd ChannelCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	chatInst := ctx.Chat()
	usr := ctx.User()
	chann := ctx.Channel()

	name := RemoveCmd(fields[0], cmdLine.String())
	if len(name) == 0 {
		io.WriteString(usr, "You need to specify a channel...\r\n")
		return
	}

	newChan, ok := chatInst.Channels()[name]
	if !ok {
		newChan = channel.NewChannel(name)
		chatInst.AddChannel(newChan)
	}
	io.WriteString(usr, fmt.Sprintf("Joined channel %[1]s\r\n", name))

	msg := chann.AddNotification(usr, fmt.Sprintf("User %[1]s left channel\r\n", usr.Name()))
	chann.SendMessage(usr, msg)
	chann.DelUser(usr)
	newChan.AddUser(usr)
	chann = newChan
	msg = chann.AddNotification(usr, fmt.Sprintf("User %[1]s joined channel\r\n", usr.Name()))
	chann.SendMessage(usr, msg)
	ctx.SetAttribute("channel", chann)
}

func init() {
	GetRegistry().RegisterCmd(ChannelCommand{})
}
