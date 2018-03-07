package commands

import (
	"io"
	"testing"

	"github.com/diltram/go-telnet"
	"github.com/diltram/go-telnet/telsh"
	"github.com/reiver/go-oi"
)

type TestCommand struct{}

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

	expCount := len(reg1.commands) + 1
	reg1.RegisterCmd(TestCommand{})

	cmdsCount := len(reg1.commands)
	if cmdsCount != expCount {
		t.Errorf("Bad amount of commands registered. Got: %d, expect: %d.", cmdsCount, expCount)
	}
}

func TestGetAllCmds(t *testing.T) {
	reg1 := GetRegistry()

	cmdsCount := len(reg1.GetAllCmds())
	expCount := len(registry.commands)

	if cmdsCount != expCount {
		t.Errorf("Bad amount of commands registered. Got: %d, expect: %d.", cmdsCount, expCount)
	}
}
