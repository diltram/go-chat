package command

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/context"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

var update = flag.Bool("update", false, "update .golden files")

func TestCallHelp(t *testing.T) {
	golden := filepath.Join("testdata", "help.golden")
	cmd := HelpCommand{}
	ctx := usrctx.NewUserContext()

	chat := chat.NewChat()
	ctx.SetAttribute("chat", chat)

	usr, server, closer := user.MockUser()
	ctx.SetAttribute("user", usr)
	defer closer()

	readBuf := new(bytes.Buffer)
	go func() {
		buf := make([]byte, 20)
		for {
			server.Read(buf)
			readBuf.Write(buf)
		}
	}()

	cmd.Call(ctx, []string{"/help"}, bytes.NewBufferString("/help"))
	actual := readBuf.Bytes()
	if *update {
		ioutil.WriteFile(golden, actual, 0644)
	}

	expected, _ := ioutil.ReadFile(golden)
	if !bytes.Equal(actual, expected) {
		t.Error("Help returned different text")
	}
}
