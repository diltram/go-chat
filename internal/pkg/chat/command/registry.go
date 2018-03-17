package command

import (
	"sync"
)

var registry *Registry

type Registry struct {
	commands   map[string]Command
	defCommand Command
	mutex      sync.RWMutex
}

func (r *Registry) RegisterCmd(cmd Command) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, key := range cmd.Cmds() {
		r.commands[key] = cmd
	}
}

func (r *Registry) RegisterDefCmd(cmd Command) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.defCommand = cmd
}

func (r *Registry) Command(key string) Command {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	cmd, ok := r.commands[key]
	if !ok {
		return r.defCommand
	}

	return cmd
}

func (r *Registry) Commands() map[string]Command {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.commands
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
		cmdsMap := make(map[string]Command)
		registry = &Registry{commands: cmdsMap}
	}

	return registry
}
