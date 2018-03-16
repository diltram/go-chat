// @TODO: Add better testing with server created and multiple users
// connected.
package channel

import (
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

	fn := func(conn net.Conn, usr *user.User) {
		called = usr.Name()
	}

	chat.Call(fn)
	if called != expected {
		t.Errorf("Call: expected %s, actual %s", expected, called)
	}
}
