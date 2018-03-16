//@TODO: Sending just a space crashes the server

package handler

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/user"
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

func (h *ChatHandler) Serve(ctx context.Context, writer io.Writer, reader io.Reader) {
	buf := make([]byte, 50)
	cmdLine := new(bytes.Buffer)

	attr, err := ctx.Attribute("chat")
	if err != nil {
		panic(err)
	}

	chatInst, ok := attr.(*chat.Chat)
	if ok == false {
		panic("Chat instance is not available")
	}

	attr, err = ctx.Attribute("user")
	if err != nil {
		panic(err)
	}

	usr, ok := attr.(*user.User)
	if ok == false {
		panic("User instance is not available")
	}

	chann := chatInst.Channels()["default"]

	// Send notification about new user
	msg := chann.AddNotification(usr, fmt.Sprintf("User %s connected to channel %s\r\n", usr.Name(), chann.Name()))
	chann.SendMessage(usr, msg)

	// Say hello to new user
	io.WriteString(writer, chatInst.WelcomeMessage())
	io.WriteString(writer, fmt.Sprintf("Nick %s, welcome in a channel %s. There is %d other users\r\n", usr.Name(), chann.Name(), len(chann.Users())-1))

	// Main handler loop
	for {
		n, err := reader.Read(buf)
		//log.Debugf("Received new command: %q", buf)

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
				io.WriteString(writer, chatInst.ExitMessage())
				break
			} else if fields[0] == "/nick" {
				// Change nick of user
				oldNick := usr.Name()
				usr.SetName(strings.Join(fields[1:], " "))
				msg := chann.AddNotification(usr, fmt.Sprintf("User %s changed his nick to %s\r\n", oldNick, usr.Name()))
				chann.SendMessage(usr, msg)
			} else {
				// Let's send message to other users
				msg := chann.AddMessage(usr, cmdLine.String())
				chann.SendMessage(usr, msg)
			}
		}

		cmdLine.Reset()
		// Write nickname of user. Disabled as looks crappy :/
		//io.WriteString(writer, fmt.Sprintf("%s: ", usr.Name()))
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
		line.Write(buf[:writeTo+2])
	} else {
		line.Write(buf)
	}
}
