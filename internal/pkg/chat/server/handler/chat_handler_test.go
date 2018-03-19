package handler

import (
	"testing"

	"github.com/diltram/go-chat/internal/pkg/server/context"
)

func TestServeBadContext(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	handler := NewChatHandler()
	ctx := new(context.Context)
	handler.Serve(*ctx)
}
