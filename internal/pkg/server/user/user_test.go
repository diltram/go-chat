package user

import (
	"net"
	"testing"
)

func TestGetters(t *testing.T) {
	tests := []struct {
		name         string // user nickname
		expectedName string // nickname from getter
	}{
		// @TODO: enable that test case when will have idea how to test random.
		//{"", "user0"},
		{"user", "user"},
		{"___", "___"},
		{"(%$*&#(*$", "(%$*&#(*$"},
	}

	for i, tt := range tests {
		conn, _ := net.Dial("tcp", ":80")
		user := NewUser(conn, tt.name)

		if user.Conn() != conn {
			t.Errorf("Getters(%d): expected %d, actual %d", i, conn, user.Conn())
		}

		if user.Name() != tt.expectedName {
			t.Errorf("Getters(%d): expected %s, actual %s", i, tt.expectedName, user.Name())
		}
	}
}

func TestSetName(t *testing.T) {
	tests := []struct {
		name string // user nickname
	}{
		{"user0"},
		{"___"},
		{"(%$*&#(*$"},
	}

	for i, tt := range tests {
		conn, _ := net.Dial("tcp", ":80")
		user := NewUser(conn, "")
		user.SetName(tt.name)

		if user.Name() != tt.name {
			t.Errorf("Setters(%d): expected %s, actual %s", i, tt.name, user.Name())
		}
	}
}
