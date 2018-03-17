package handler

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/server/context"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

// EchoHandler implements base handler which echoes everything send to it.
type EchoHandler struct{}

// Serve handler and echo for all requests.
func (h EchoHandler) Serve(ctx context.Context) {
	attr, err := ctx.Attribute("user")
	if err != nil {
		panic(err)
	}

	usr, ok := attr.(*user.User)
	if ok == false {
		panic("User instance is not available")
	}

	buf := make([]byte, 8)
	for {
		if _, err := io.CopyBuffer(usr, usr, buf); err != nil {
			log.Fatal(err)
			break
		}
	}
}
