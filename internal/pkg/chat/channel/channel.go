package channel

import (
	"io"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/pkg/chat/message"
	"github.com/diltram/go-chat/internal/pkg/server/user"
)

const (
	DefaultChannelName = "default"
)

// Call represents a function which can be provided to Call method.
// When iterating over all users it can execute some specific operation e.g.
// send new message.
type Call func(net.Conn, *user.User)

// Channel provides base functionality of a server.
// It keeps track of messages sent and all users currently connected to server.
type Channel struct {
	name     string                  // Name of the chat
	users    map[net.Conn]*user.User // Map of all users connected to the chat
	messages []message.Stringer      // Slice with messages sent to chat
	mutex    sync.RWMutex            // mutex for access to messages and users
}

// NewChannel creates instance of new chat with specified name.
func NewChannel(name string) *Channel {
	users := make(map[net.Conn]*user.User)
	messages := make([]message.Stringer, 5, 5)
	return &Channel{
		name:     name,
		users:    users,
		messages: messages,
	}
}

func (c Channel) Name() string {
	return c.name
}

// AddUser registers new user into specific chat.
func (c *Channel) AddUser(user *user.User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.users[user.Conn()] = user
}

// DelUser de-registers user from chat.
func (c *Channel) DelUser(usr *user.User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.users, usr.Conn())
}

// Users provides access to information about all users in chat.
func (c *Channel) Users() map[net.Conn]*user.User {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.users
}

// Call allows to iterate over all users registered in chat and call them.
// That method allows for, for a example sending new messages to all users.
func (c *Channel) Call(fn Call) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	log.Debugf("Executing call on %d users", len(c.users))
	for conn, user := range c.users {
		fn(conn, user)
	}
}

// AddMessage registers new message in chat.
func (c *Channel) AddMessage(sender *user.User, content string) message.Stringer {
	msg := message.NewMessage(sender, content, time.Now())
	c.addStringer(msg)

	return msg
}

// AddNotification registers new notification in chat.
func (c *Channel) AddNotification(sender *user.User, content string) message.Stringer {
	msg := message.NewNotification(sender, content, time.Now())
	c.addStringer(msg)

	return msg
}

func (c *Channel) addStringer(msg message.Stringer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.messages = append(c.messages, msg)
}

func (c *Channel) SendMessage(sender *user.User, msg message.Stringer) {
	c.Call(func(n net.Conn, u *user.User) {
		if u != sender {
			w := n.(io.Writer)
			io.WriteString(w, msg.String())
		}
	})
}
