package message

import (
	"fmt"
	"time"

	"github.com/diltram/go-chat/internal/pkg/server/user"
)

type Notification struct {
	date   time.Time  // date/time when message was sent
	text   string     // content of the message
	sender *user.User // user which sent message
}

// String returns textual representation of a message.
// It includes all information like date when message was sent and who send
// message.
func (n *Notification) String() string {
	return fmt.Sprintf("%s | %s\r\n", n.date.Format("03/2/2006 15:04:05"), n.text)
}

// NewMessage generates a message with information about the user, content and
// time.
func NewNotification(sender *user.User, content string, date time.Time) Stringer {
	return &Notification{
		date:   date,
		text:   content,
		sender: sender,
	}
}
