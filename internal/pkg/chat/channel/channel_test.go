// @TODO: Add better testing with server created and multiple users
// connected.
package channel

import (
	"bytes"
	"net"
	"reflect"
	"testing"

	"github.com/diltram/go-chat/internal/pkg/server/user"
)

func TestAddDelUsers(t *testing.T) {
	channel := NewChannel("First channel")

	conn1, _ := net.Dial("tcp", ":80")
	user1 := user.NewUser(conn1, "user")
	channel.AddUser(user1)

	expected := map[net.Conn]*user.User{
		conn1: user1,
	}

	if !reflect.DeepEqual(channel.Users(), expected) {
		t.Errorf("Users: expected %v, actual %v", expected, channel.Users())
	}

	channel.DelUser(user1)
	expected = map[net.Conn]*user.User{}

	if !reflect.DeepEqual(channel.Users(), expected) {
		t.Errorf("Users: expected %v, actual %v", expected, channel.Users())
	}
}

func TestCall(t *testing.T) {
	chat := NewChannel("First chat")
	var called string
	expected := "user"

	conn1, _ := net.Dial("tcp", ":80")
	user1 := user.NewUser(conn1, expected)
	chat.AddUser(user1)

	fn := func(usr *user.User) {
		called = usr.Name()
	}

	chat.Call(fn)
	if called != expected {
		t.Errorf("Call: expected %s, actual %s", expected, called)
	}
}

func TestAddMessage(t *testing.T) {
	chat := NewChannel("First chat")
	conn1, _ := net.Dial("tcp", ":80")
	user1 := user.NewUser(conn1, "")
	chat.AddUser(user1)
	expected := len(chat.messages) + 1

	message := "test content"
	msg := chat.AddMessage(user1, message)

	actual := len(chat.messages)
	if actual != expected {
		t.Errorf("AddMessage: expected %d, actual %d", expected, actual)
	}

	actualMsg := chat.messages[len(chat.messages)-1]
	if actualMsg != msg {
		t.Errorf("AddMessage: expecting %+v, actual %+v", msg, actualMsg)
	}
}

func TestAddNotification(t *testing.T) {
	chat := NewChannel("First chat")
	conn1, _ := net.Dial("tcp", ":80")
	user1 := user.NewUser(conn1, "")
	chat.AddUser(user1)
	expected := len(chat.messages) + 1

	message := "test content"
	msg := chat.AddNotification(user1, message)

	actual := len(chat.messages)
	if actual != expected {
		t.Errorf("AddNotification: expected %d, actual %d", expected, actual)
	}

	actualMsg := chat.messages[len(chat.messages)-1]
	if actualMsg != msg {
		t.Errorf("AddNotification: expecting %+v, actual %+v", msg, actualMsg)
	}
}

func TestSendMessage(t *testing.T) {
	chat := NewChannel("First chat")
	user1, _, closer1 := user.MockUser()
	user2, server2, closer2 := user.MockUser()
	defer closer1()
	defer closer2()

	chat.AddUser(user1)
	chat.AddUser(user2)

	readBuf := new(bytes.Buffer)
	go func() {
		buf := make([]byte, 1)
		for {
			server2.Read(buf)
			readBuf.Write(buf)
		}
	}()

	readBuf.Reset()
	message := "test content"
	msg := chat.AddMessage(user1, message)
	chat.SendMessage(user1, msg)

	actual := readBuf.String()
	expected := msg.String()
	if actual != expected {
		t.Errorf("AddNotification: expected %s, actual \"%s\"", expected, actual)
	}
}
