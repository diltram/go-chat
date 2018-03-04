package app

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/diltram/go-chat/internal/app/go-chat/commands"
	"github.com/diltram/go-chat/internal/app/go-chat/config"
)

const (
	appName  = "go-chat"
	appUsage = "Golang implementation of telnet chat server"
)

var app *cli.App

func init() {
	var conf config.Configuration
	app = cli.NewApp()
	app.Version = "Darwin.1.00.0"
	app.Name = appName
	app.Usage = appUsage
	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Printf("unknown command - %v \n\n", command)
		cli.ShowAppHelp(ctx)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "./configs/config.yaml",
			Usage:  "load configuration from `FILE`",
			EnvVar: "GO_CHAT_CONFIG",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug logs",
			EnvVar: "GO_CHAT_DEBUG",
		},
		cli.StringFlag{
			Name:   "log-format",
			Value:  "text",
			Usage:  "set the format used by logs ('text', or 'json')",
			EnvVar: "GO_CHAT_LOG_FORMAT",
		},
	}

	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		switch ctx.GlobalString("log-format") {
		case "text":
		case "json":
			log.SetFormatter(new(log.JSONFormatter))
		default:
			log.Fatalf("unknown log-format %q", ctx.GlobalString("log-format"))
		}

		conf = config.LoadConfig(ctx.GlobalString("config"))

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Starts telnet server",
			Action: func(ctx *cli.Context) error {
				return commands.OnServer(conf)
			},
		},
	}
}

func runCli() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
