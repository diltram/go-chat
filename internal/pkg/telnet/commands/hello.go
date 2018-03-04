package commands

import (
	"io"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
)

//HelloCmd provides /hello in telnet
type HelloCmd struct{}

//Name provides a name under which command will be available in telnet
func (cmd HelloCmd) Name() string {
	return "/hello"
}

//Description provides a long description which will be used for help in telnet
func (cmd HelloCmd) Description() string {
	return ""
}

//Produce exposes command in format required by telnet library
func (cmd HelloCmd) Produce() telsh.ProducerFunc {
	return telsh.ProducerFunc(cmd.cmdProducer)
}

func (cmd HelloCmd) cmdHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "  Hello!")

	return nil
}

func (cmd HelloCmd) cmdProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(cmd.cmdHandler)
}

func init() {
	GetRegistry().RegisterCmd(new(HelloCmd))
}
