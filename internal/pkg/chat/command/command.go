package command

import (
	"bytes"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

type Command interface {
	Name() string
	Desc() string
	Cmds() []string
	Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer)
}
