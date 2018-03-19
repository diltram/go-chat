package command

import (
	"bytes"

	"github.com/diltram/go-chat/internal/pkg/chat/context"
)

// Command specifies how command structure need to looks like.
// It specifies all the methods as they're all required by help command.
type Command interface {
	Name() string
	Desc() string
	Cmds() []string
	Call(ctx *usrctx.UserContext, fields []string, cmdLine *bytes.Buffer)
}
