package commands

import (
	"io"
	"testing"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
)

type TestCommand struct {}

func (cmd TestCommand) Name() string {
	return "/test"
}

func (cmd TestCommand) Description() string {
	return ""
}

func (cmd TestCommand) Produce() telsh.ProducerFunc {
	return telsh.ProducerFunc(cmd.cmdProducer)
}

func (cmd TestCommand) cmdHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "Test!\r\n")
	return nil
}

func (cmd TestCommand) cmdProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(cmd.cmdHandler)
}

func TestGetRegistry(t *testing.T) {
	reg1 := GetRegistry()
	reg2 := GetRegistry()

	if reg1 != reg2 {
		t.Errorf("Got different registry. Singleton doesn't work.")
	}
}

func TestRegisterCmd(t *testing.T) {
	reg1 := GetRegistry()

	exp_cmds_count := len(reg1.commands) + 1
	reg1.RegisterCmd(TestCommand{})

	cmds_count := len(reg1.commands)
	if cmds_count != exp_cmds_count {
		t.Errorf("Bad amount of commands registered. Got: %d, expect: %d.", cmds_count, exp_cmds_count)
	}
}

func TestGetAllCmds(t *testing.T) {
	reg1 := GetRegistry()

	cmds_count := len(reg1.GetAllCmds())
	exp_count := len(registry.commands)

	if cmds_count != exp_count {
		t.Errorf("Bad amount of commands registered. Got: %d, expect: %d.", cmds_count, exp_count)
	}
}
