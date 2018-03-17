package command

import (
	"bytes"
	"fmt"
	"io"
	"sort"
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
	unqCmds := make(map[string]Command)
	unqNames := make([]string, 3)
	usr := ctx.User()
	commands := GetRegistry().Commands()

	io.WriteString(usr, "\r\nHelp:\r\n")
	io.WriteString(usr, "--------------------\r\n")

	for _, v := range commands {
		if _, ok := unqCmds[v.Name()]; !ok {
			unqNames = append(unqNames, v.Name())
			unqCmds[v.Name()] = v
		}
	}

	sort.Strings(unqNames)
	for _, name := range unqNames {
		cmd, ok := unqCmds[name]
		if ok {
			line := fmt.Sprintf("%20s %15s %s\r\n", strings.Join(cmd.Cmds(), ", "), cmd.Name(), cmd.Desc())
			io.WriteString(usr, line)
		}
	}
}

func init() {
	GetRegistry().RegisterCmd(HelpCommand{})
}
