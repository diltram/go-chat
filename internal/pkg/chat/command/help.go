package command

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// HelpCommand prints help for all registered commands.
// It loads commands from commands registry so whenever new command will be
// created and registered it will be automatically updated here.
type HelpCommand struct{}

// Name informs that command prints help.
func (cmd HelpCommand) Name() string {
	return "Help"
}

// Desc returns longer description of the command.
func (cmd HelpCommand) Desc() string {
	return "Print out that help"
}

// Cmds returns a list of names which can be used to access help.
func (cmd HelpCommand) Cmds() []string {
	return []string{"/help", "/h"}
}

// Call loads all registered commands from commands registry.
// After sorting them by name it prints names, descriptions and access names of
// all commands.
func (cmd HelpCommand) Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer) {
	usr := ctx.User()
	commands := GetRegistry().UniqueCommands()

	io.WriteString(usr, "\r\nHelp:\r\n")
	io.WriteString(usr, "--------------------\r\n")
	for _, cmd := range commands {
		line := fmt.Sprintf("%20s %15s %s\r\n", strings.Join(cmd.Cmds(), ", "), cmd.Name(), cmd.Desc())
		io.WriteString(usr, line)
	}
}

func init() {
	GetRegistry().RegisterCmd(HelpCommand{})
}
