package message

import (
	"fmt"
	"time"

	"github.com/diltram/go-chat/internal/pkg/server/user"
)

// Message represents one message sent by user on the chat.
type Message struct {
	date   time.Time  // date/time when message was sent
	text   string     // content of the message
	sender *user.User // user which sent message
}

// String returns textual representation of a message.
// It includes all information like date when message was sent and who send
// message.
func (m *Message) String() string {
	return fmt.Sprintf("%s | %s - %s", m.date.Format("03/2/2006 15:04:05"), m.sender.Name(), m.text)
}

// NewMessage generates a message with information about the user, content and
// time.
func NewMessage(sender *user.User, content string, date time.Time) Stringer {
	return &Message{
		date:   date,
		text:   content,
		sender: sender,
	}
}
