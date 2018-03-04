package commands

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type FakeWriter struct {
	io.WriteCloser

	LastWrite string
}

func (fk FakeWriter) Write(p []byte) (n int, err error) {
	fk.LastWrite = string(p)
	return 0, nil
}

func TestProduce(t *testing.T) {
	cmd := HelloCmd{}

	if cmd.Produce() == nil {
		t.Error("Produce returned nil, should return function")
	}
}

func TestCmdHandler(t *testing.T) {
	cmd := HelloCmd{}
	writer := new(FakeWriter)
	cmd.cmdHandler(ioutil.NopCloser(strings.NewReader("")), writer, new(FakeWriter))

	expected := "  Hello!"
	if writer.LastWrite != expected {
		t.Errorf("Wrong message returned. Expecting: %s, received: %s", expected, writer.LastWrite)
	}
}
