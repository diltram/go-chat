package handler

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/chat/command"
	"github.com/diltram/go-chat/internal/pkg/chat/context"
	"github.com/diltram/go-chat/internal/pkg/server/context"
	"github.com/diltram/go-chat/internal/pkg/server/handler"
)

const (
	defaultExitCommandName = "/quit"
	lineEnd                = "\r\n"
)

type ChatHandler struct {
	ExitCommandName string
}

func NewChatHandler() handler.Handler {
	return &ChatHandler{
		ExitCommandName: defaultExitCommandName,
	}
}

func (h *ChatHandler) Serve(ctx context.Context) {
	buf := make([]byte, 50)
	cmdLine := new(bytes.Buffer)

	usrCtx, ok := ctx.(*usrctx.UserContext)
	if !ok {
		panic("Provided context is wrong. Should be UserContext")
	}

	chatInst := usrCtx.Chat()
	usr := usrCtx.User()
	chann := usrCtx.Channel()

	// Send notification about new user
	msg := chann.AddNotification(usr, fmt.Sprintf("User %s connected to channel %s", usr.Name(), chann.Name()))
	chann.SendMessage(usr, msg)

	// Say hello to new user
	io.WriteString(usr, chatInst.WelcomeMessage())
	io.WriteString(usr, fmt.Sprintf("Nick %s, welcome in a channel %s. There is %d other users\r\n", usr.Name(), chann.Name(), len(chann.Users())-1))

	// Main handler loop
	for {
		n, err := usr.Read(buf)
		if n <= 0 && err == nil {
			// Nothing to do, let's check again
			continue
		} else if n <= 0 && err != nil {
			// Some error happended and not connection closed. Log it and end.
			if err != io.EOF {
				log.Error(err)
			}
			break
		}
		buf = bytes.Trim(buf, " ")
		h.writeBuf(cmdLine, buf)

		// Check if the line is ended
		if h.isCompleteLine(buf) {
			if h.isEmptyLine(cmdLine) {
				h.clear(cmdLine)
				continue
			}

			fields := strings.Fields(cmdLine.String())
			log.Debugf("Found %d tokens: %v", len(fields), fields)

			if fields[0] == h.ExitCommandName {
				io.WriteString(usr, chatInst.ExitMessage())
				break
			}

			cmd := command.GetRegistry().Command(fields[0])
			cmd.Call(usrCtx, fields, cmdLine)
		}
		cmdLine.Reset()
	}
}

func (h ChatHandler) isEmptyLine(line *bytes.Buffer) bool {
	lineEndOnly := bytes.Equal(line.Bytes()[:2], []byte(lineEnd))

	if line.Len() > 2 && lineEndOnly {
		return true
	}

	return false
}

func (h ChatHandler) clear(line *bytes.Buffer) {
	line.Reset()
}

// isCompleteLine checks if there is \n in text send.
// Checks that on buffer to speed up processing and don't go over growing line.
func (h ChatHandler) isCompleteLine(buf []byte) bool {
	if bytes.Index(buf, []byte(lineEnd)) > -1 {
		return true
	}

	return false
}

func (h ChatHandler) writeBuf(line *bytes.Buffer, buf []byte) {
	writeTo := bytes.Index(buf, []byte(lineEnd))

	if writeTo > 0 {
		line.Write(buf[:writeTo])
	} else {
		line.Write(buf)
	}
}
