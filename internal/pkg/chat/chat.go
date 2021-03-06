package chat

import (
	"sync"

	"github.com/diltram/go-chat/internal/pkg/chat/channel"
)

// Chat handles information about all the channels and provides base data like
// welcome and exit messages.
type Chat struct {
	channels       map[string]*channel.Channel
	welcomeMessage string
	exitMessage    string
	mutex          sync.RWMutex
}

// NewChat creates new instance of a chat.
func NewChat() *Chat {
	chat := &Chat{
		channels:       make(map[string]*channel.Channel),
		welcomeMessage: "\r\nHello!\r\n",
		exitMessage:    "\r\nGoodbye!\r\n",
	}

	chann := channel.NewChannel(channel.DefaultChannelName)
	chat.AddChannel(chann)

	return chat
}

// WelcomeMessage returns message which every user will see after connecting to
// server.
func (c *Chat) WelcomeMessage() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.welcomeMessage
}

// SetWelcomeMessage allows to update a welcome message.
func (c *Chat) SetWelcomeMessage(msg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.welcomeMessage = msg
}

// ExitMessage returns string which will be presented to user when he's leaving
// the chat.
func (c *Chat) ExitMessage() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.exitMessage
}

// SetExitMessage updates a message with new text.
func (c *Chat) SetExitMessage(msg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.exitMessage = msg
}

// AddChannel registers new channel on the server.
func (c *Chat) AddChannel(chann *channel.Channel) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.channels[chann.Name()] = chann
}

// Channels returns a map of all channels registered on the server.
func (c *Chat) Channels() map[string]*channel.Channel {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.channels
}
