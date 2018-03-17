package command

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// ChannelCommand allows to change a channel.
type ChannelCommand struct{}

func (cmd ChannelCommand) Name() string {
	return "Change channel"
}

func (cmd ChannelCommand) Desc() string {
	return "Commands allow to change current channel into a new one."
}

func (cmd ChannelCommand) Cmds() []string {
	return []string{"/channel", "/chan", "/ch"}
}

func (cmd ChannelCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	chatInst := ctx.Chat()
	usr := ctx.User()
	chann := ctx.Channel()

	name := strings.TrimLeft(cmdLine.String(), fields[0]+" ")
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
