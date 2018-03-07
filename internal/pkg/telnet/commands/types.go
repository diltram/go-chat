package commands

import (
	"github.com/diltram/go-telnet/telsh"
)

// CommandsRegistry is a singleton for keeping information about all commands in Telnet.
type CommandsRegistry struct {
	commands []TelnetCommand
}

// TelnetCommand is interface for all Telnet commands.
type TelnetCommand interface {
	Name() string
	Description() string
	Produce() telsh.ProducerFunc
}
