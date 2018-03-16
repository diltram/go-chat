package handler

import (
	"io"

	"github.com/diltram/go-chat/internal/pkg/server/context"
)

// Handler is interface that wraps the Serve method.
type Handler interface {
	Serve(ctx context.Context, writer io.Writer, reader io.Reader)
}
