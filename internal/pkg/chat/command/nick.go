package command

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// NickCommand allows to change a nick.
type NickCommand struct{}

func (cmd NickCommand) Name() string {
	return "Change nick"
}

func (cmd NickCommand) Desc() string {
	return "Commands allow to change current nick into a new one."
}

func (cmd NickCommand) Cmds() []string {
	return []string{"/n", "/nick"}
}

func (cmd NickCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	usr := ctx.User()
	chann := ctx.Channel()

	nick := strings.TrimLeft(cmdLine.String(), fields[0]+" ")
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
