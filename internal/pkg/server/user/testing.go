package user

import (
	"net"
)

// MockUser function creates client/server connection and returns User
// configured with it.
//
// It returns User, server and closer function which closes server and client.
// Before ending a test you should call it to cleanup session.
func MockUser() (*User, net.Conn, func()) {
	server, client := net.Pipe()
	closer := func() {
		server.Close()
		client.Close()
	}
	user := NewUser(client, "")

	return user, server, closer
}
