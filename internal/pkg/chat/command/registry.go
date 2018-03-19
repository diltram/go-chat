package command

import (
	"sort"
	"sync"
)

var registry *Registry

// Registry stores information about all commands available in chat.
// When any new command is created it should be registered in registry using
// RegisterCmd method.
type Registry struct {
	commands   map[string]Command
	defCommand Command
	mutex      sync.RWMutex
}

// RegisterCmd registers new command in registry.
// It will register command under all of it's command names for example nick
// command is registered under /nick and /n commands.
// It's possible to override command name. There will be no error produced when
// it will happen.
func (r *Registry) RegisterCmd(cmd Command) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, key := range cmd.Cmds() {
		r.commands[key] = cmd
	}
}

// RegisterDefCmd shouldn't be used.
// It's used to register MsgCommand and without need shouldn't be change.
func (r *Registry) RegisterDefCmd(cmd Command) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.defCommand = cmd
}

// Command returns a command based on the name.
// When there is no command registered under that name it will return default
// MsgCommand.
func (r *Registry) Command(key string) Command {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	cmd, ok := r.commands[key]
	if !ok {
		return r.defCommand
	}

	return cmd
}

// Commands provides access to all of the registered commands.
func (r *Registry) Commands() map[string]Command {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.commands
}

// UniqueCommands provides a slice of sorted commands which are registered.
// The commands are sorted based on the name.
// It doesn't matter under how many names command is registered. It will be
// provided by that method only once.
func (r *Registry) UniqueCommands() []Command {
	unqNames := make(map[string]bool)
	unqCmds := make([]Command, 0, 5)

	for _, v := range r.Commands() {
		if _, ok := unqNames[v.Name()]; !ok {
			unqNames[v.Name()] = true
			unqCmds = append(unqCmds, v)
		}
	}

	sort.Slice(unqCmds, func(i, j int) bool {
		return unqCmds[i].Name() < unqCmds[j].Name()
	})

	return unqCmds
}

// NewRegistry creates empty registry.
func NewRegistry() *Registry {
	cmdsMap := make(map[string]Command)
	return &Registry{
		commands: cmdsMap,
	}
}

// GetRegistry returns Registry singleton.
//
// That method is not thread safe so it's possible to have race conditions with
// two threads creating new registry in same time. Because of that one command
// can be unavailabe.
// That part of code should be single threaded in my app so leaving that as is
// for now.
func GetRegistry() *Registry {
	if registry == nil {
		registry = NewRegistry()
	}

	return registry
}
