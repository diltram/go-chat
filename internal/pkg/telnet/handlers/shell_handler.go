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

func (handler *ShellHandlerCommands) addClient(c net.Conn) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	handler.clients[c] = handler.clientsCount
	log.Debug("Registered new user #", handler.clientsCount)
	handler.clientsCount++
}

func (handler *ShellHandlerCommands) removeClient(c net.Conn) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	log.Debug("Closing connection with user ", handler.clients[c])
	delete(handler.clients, c)
}

func (handler *ShellHandlerCommands) ServeTELNET(ctx telnet.Context,
	writer telnet.Writer,
	reader telnet.Reader) {

	handler.addClient(ctx.Conn())
	defer handler.removeClient(ctx.Conn())

	handler.ShellHandler.ServeTELNET(ctx, writer, reader)
}
