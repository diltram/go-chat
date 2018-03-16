package handler

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/server/context"
)

// EchoHandler implements base handler which echoes everything send to it.
type EchoHandler struct{}

// Serve handler and echo for all requests.
func (h EchoHandler) Serve(ctx context.Context, w io.Writer, r io.Reader) {
	buf := make([]byte, 8)
	for {
		if _, err := io.CopyBuffer(w, r, buf); err != nil {
			log.Fatal(err)
			break
		}
	}
}
