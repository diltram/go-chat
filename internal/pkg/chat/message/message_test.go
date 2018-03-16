package message

import (
	"net"
	"testing"
	"time"

	"github.com/diltram/go-chat/internal/pkg/chat/user"
)

func TestString(t *testing.T) {
	tests := []struct {
		name      string    // message sender
		content   string    // message content
		date      time.Time // when message was sent
		expString string    // expected stringified version of message
	}{
		{
			"user",
			"Hey!!!",
			time.Date(2018, time.March, 18, 10, 0, 0, 0, time.UTC),
			"10/18/2018 10:00:00 | user - Hey!!!",
		},
		{
			"john",
			"How are you brother?",
			time.Date(2018, time.January, 1, 1, 0, 0, 0, time.UTC),
			"01/1/2018 01:00:00 | john - How are you brother?",
		},
	}

	for i, tt := range tests {
		conn, _ := net.Dial("tcp", ":80")
		sender := user.NewUser(conn, tt.name)
		msg := NewMessage(sender, tt.content, tt.date)

		if msg.String() != tt.expString {
			t.Errorf("String(%d): expected %s, actual %s", i, tt.expString, msg.String())
		}
	}
}
