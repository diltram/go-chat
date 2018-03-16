package user

import (
	"fmt"
	"math/rand"
	"net"
)

// User represents data about specific connection.
//
// It keeps information about the name of the user, his connection and any
// other required data.
type User struct {
	conn net.Conn // connection
	name string   // nickname
}

// Conn returns user's connection.
func (u *User) Conn() net.Conn {
	return u.conn
}

// Name returns user's nickname.
func (u *User) Name() string {
	return u.name
}

// SetName allows to change user's nickname.
func (u *User) SetName(name string) {
	u.name = name
}

// NewUser creates a new user with specified connection and nickname.
func NewUser(conn net.Conn, name string) *User {
	if "" == name {
		name = fmt.Sprintf("user%d", rand.Intn(1000000))
	}

	return &User{conn: conn, name: name}
}
