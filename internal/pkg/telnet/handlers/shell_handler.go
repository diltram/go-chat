package handlers

import (
	"net"
	"sync"

	"github.com/diltram/go-telnet"
	"github.com/diltram/go-telnet/telsh"
	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/telnet/commands"
)

type ShellHandlerCommands struct {
	*telsh.ShellHandler
	clientsCount int
	clients      map[net.Conn]int
	mutex        sync.Mutex
}

//NewShellHandler creates new handler for telnet server.
//All available commands will be automatically registered so no others steps
//required.
func NewShellHandler() *ShellHandlerCommands {
	telnetHandler := ShellHandlerCommands{
		ShellHandler: telsh.NewShellHandler(),
		clientsCount: 0,
		clients:      make(map[net.Conn]int),
		mutex:        sync.Mutex{},
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
