package chat

import (
	"testing"
)

func TestSetGetWelcomeMessage(t *testing.T) {
	inst := NewChat()
	msg := "Some completely new message\r\n"
	inst.SetWelcomeMessage(msg)

	actual := inst.WelcomeMessage()

	if actual != msg {
		t.Errorf("WelcomeMessage. Expected %s, actual %s", msg, actual)
	}
}

func TestSetGetExitMessage(t *testing.T) {
	inst := NewChat()
	msg := "Some completely new message\r\n"
	inst.SetExitMessage(msg)

	actual := inst.ExitMessage()

	if actual != msg {
		t.Errorf("ExitMessage. Expected %s, actual %s", msg, actual)
	}
}

func TestNewChannel(t *testing.T) {
	inst := NewChat()
	channels := inst.Channels()

	if len(channels) != 1 {
		t.Error("Newly created channel should have default channel only")
	}

	expected := channels["default"]
	if expected.Name() != "default" {
		t.Errorf("First channel should be named default, actual name is %s", expected.Name())
	}
}
