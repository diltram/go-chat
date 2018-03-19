package command

import (
	"testing"
)

func TestRegisteredCommands(t *testing.T) {
	cmds := GetRegistry().UniqueCommands()

	for _, cmd := range cmds {
		if cmd.Name() == "" {
			t.Errorf("%T, expecting non-empty name", cmd)
		}

		if cmd.Desc() == "" {
			t.Errorf("%T, expecting non-empty description", cmd)
		}

		if len(cmd.Cmds()) == 0 {
			t.Errorf("%T, command need to have at least one cmd name to call it", cmd)
		}
	}
}

func TestNewRegistry(t *testing.T) {
	cmdsRegistry := NewRegistry()

	if len(cmdsRegistry.Commands()) != 0 {
		t.Error("Newly created registry should be empty")
	}

	if len(cmdsRegistry.UniqueCommands()) != 0 {
		t.Error("Newly created registry should be empty")
	}
}

func TestCommandGetterDefault(t *testing.T) {
	msgCmd := MsgCommand{}
	cmdsRegistry := NewRegistry()
	cmdsRegistry.RegisterDefCmd(msgCmd)

	if len(cmdsRegistry.Commands()) != 0 {
		t.Error("Default command should not be registered into standard list")
	}

	cmd := cmdsRegistry.Command("example")
	if cmd != msgCmd {
		t.Errorf("Expecting default command to be %T, actual %T", msgCmd, cmd)
	}
}

func TestCommandGetter(t *testing.T) {
	channelCmd := ChannelCommand{}
	helpCmd := HelpCommand{}
	cmdsRegistry := NewRegistry()

	cmdsRegistry.RegisterCmd(helpCmd)
	cmdsRegistry.RegisterCmd(channelCmd)
	if len(cmdsRegistry.Commands()) == 0 {
		t.Error("Commands should be registered into commands map")
	}

	cmd := cmdsRegistry.Command(channelCmd.Cmds()[0])
	if cmd != channelCmd {
		t.Errorf("Expecting command to be %T, actual %T", channelCmd, cmd)
	}

	// Test that all cmds names are registered
	cmd = cmdsRegistry.Command(channelCmd.Cmds()[1])
	if cmd != channelCmd {
		t.Errorf("Expecting command to be %T, actual %T", channelCmd, cmd)
	}
}

func TestUniqueCommands(t *testing.T) {
	channelCmd := ChannelCommand{}
	helpCmd := HelpCommand{}
	cmdsRegistry := NewRegistry()

	cmdsRegistry.RegisterCmd(channelCmd)
	cmdsRegistry.RegisterCmd(helpCmd)
	cmds := cmdsRegistry.UniqueCommands()

	expected := 2
	actual := len(cmds)
	if actual != expected {
		t.Errorf("Unique commands list should contain %d commands, actual %d", expected, actual)
	}
}
