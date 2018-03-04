package commands

import (
	"fmt"

	"github.com/reiver/go-telnet"
	log "github.com/sirupsen/logrus"

	"github.com/diltram/go-chat/internal/app/go-chat/config"
	"github.com/diltram/go-chat/internal/pkg/telnet/handlers"
)

//OnServer implements the 'server' go-chat command
func OnServer(conf config.Configuration) error {
	log.Info("Starting go-chat telnet server")

	log.Infof("Creating new telnet handler and registering all commands")
	handler := handlers.NewShellHandler()

	addr := fmt.Sprintf("%s:%d", conf.Server.IP, conf.Server.Port)
	log.Info("Starting telnet server on ", addr)

	if err := telnet.ListenAndServe(addr, handler); nil != err {
		panic(err)
	}

	return nil
}
