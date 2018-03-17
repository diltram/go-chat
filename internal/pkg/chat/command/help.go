package command

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// HelpCommand prints help with all commands listed.
type HelpCommand struct{}

// Name returns just empty string.
func (cmd HelpCommand) Name() string {
	return "Help"
}

// Desc returns just empty string.
func (cmd HelpCommand) Desc() string {
	return "Print out that help"
}

func (cmd HelpCommand) Cmds() []string {
	return []string{"/help", "/h"}
}

func (cmd HelpCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	unqCmds := make(map[Command]bool)
	usr := ctx.User()
	commands := GetRegistry().Commands()

	io.WriteString(usr, "\r\nHelp:\r\n")
	io.WriteString(usr, "--------------------\r\n")

	// Convert standard map into unique map of commands
	for _, v := range commands {
		unqCmds[v] = true
	}

	for cmd, _ := range unqCmds {
		line := fmt.Sprintf("%20s %15s %s\r\n", strings.Join(cmd.Cmds(), ", "), cmd.Name(), cmd.Desc())
		io.WriteString(usr, line)
	}
}

func init() {
	GetRegistry().RegisterCmd(HelpCommand{})
}
