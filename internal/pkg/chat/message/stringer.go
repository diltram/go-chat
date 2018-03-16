package message

// Stringer describes interface for all type of messages on the chat.
//
// Base implementation is message which shows user's message on chat.
// Another implementation to notification which differently formats content.
type Stringer interface {
	String() string
}
