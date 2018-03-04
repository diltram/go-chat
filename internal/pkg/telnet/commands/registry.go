package commands

import (
	log "github.com/sirupsen/logrus"
)

var registry *CommandsRegistry

//RegisterCmd registers new command in registry.
func (registry *CommandsRegistry) RegisterCmd(cmd TelnetCommand) {
	registry.commands = append(registry.commands, cmd)
}

//GetAllCommands returns map containing all available commands.
func (registry CommandsRegistry) GetAllCmds() []TelnetCommand {
	return registry.commands
}

//Return commands registry singleton.
func GetRegistry() *CommandsRegistry {
	if registry == nil {
		log.Debug("Creating new commands registry")
		registry = &CommandsRegistry{}
	}

	return registry
}
