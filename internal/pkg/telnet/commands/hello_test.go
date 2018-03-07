package commands

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

type MyWriteCloser struct {
	*bytes.Buffer
}

func (mwc *MyWriteCloser) Close() error {
	// Noop
	return nil
}

func TestProduce(t *testing.T) {
	cmd := HelloCmd{}

	if cmd.Produce() == nil {
		t.Error("Produce returned nil, should return function")
	}
}

func TestCmdHandler(t *testing.T) {
	cmd := HelloCmd{}
	stdin := ioutil.NopCloser(strings.NewReader(""))
	stdout := MyWriteCloser{new(bytes.Buffer)}
	stderr := MyWriteCloser{new(bytes.Buffer)}

	cmd.cmdHandler(stdin, &stdout, &stderr)

	expected := "Hello!"
	actual := stdout.Buffer.String()
	if actual != expected {
		t.Errorf("Wrong message returned. Expecting: %s, received: %s", expected, actual)
	}
}
