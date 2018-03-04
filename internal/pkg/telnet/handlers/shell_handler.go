package handlers

import (
	"github.com/reiver/go-telnet/telsh"
	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/telnet/commands"
)

type ShellHandlerCommands struct {
	*telsh.ShellHandler
}

//NewShellHandler creates new handler for telnet server.
//All available commands will be automatically registered so no others steps
//required.
func NewShellHandler() *ShellHandlerCommands {
	telnetHandler := ShellHandlerCommands{
		telsh.NewShellHandler(),
	}

	telnetHandler.registerAllCommands()

	return &telnetHandler
}

func (handler *ShellHandlerCommands) registerAllCommands() {
	commands := commands.GetRegistry().GetAllCmds()

	log.Info("Registering commands in handler")
	for _, cmd := range commands {
		log.Info("\tNew command registered:\t", cmd.Name())
		handler.Register(cmd.Name(), cmd.Produce())
	}
}
