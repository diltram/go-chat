package commands

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/app/go-chat/config"
	"github.com/diltram/go-chat/internal/pkg/chat"
	"github.com/diltram/go-chat/internal/pkg/chat/channel"
	"github.com/diltram/go-chat/internal/pkg/chat/server"
	"github.com/diltram/go-chat/internal/pkg/chat/server/handler"
	"github.com/diltram/go-chat/internal/pkg/server/context"
)

// OnServer implements the 'server' go-chat command
func OnServer(conf config.Configuration) error {
	log.Info("Starting go-chat telnet server")

	log.Infof("Creating new telnet handler and registering all commands")

	addr := fmt.Sprintf("%s:%d", conf.Server.IP, conf.Server.Port)
	log.Info("Starting telnet server on ", addr)

	chat := chat.NewChat()
	chann := channel.NewChannel(channel.DefaultChannelName)
	chat.AddChannel(chann)

	ctx := context.NewContext()
	ctx.SetAttribute("chat", chat)

	srv := server.NewServer(addr, handler.NewChatHandler(), ctx)
	if err := srv.ListenAndServe(); nil != err {
		panic(err)
	}

	return nil
}
