package app

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	appName  = "go-chat"
	appUsage = "Golang implementation of telnet chat server"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Version = "Darwin.1.00.0"
	app.Name = appName
	app.Usage = appUsage
	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Printf("unknown command - %v \n\n", command)
		cli.ShowAppHelp(ctx)
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug logs",
		},
		cli.StringFlag{
			Name:  "log",
			Usage: "log file to store logs",
		},
		cli.StringFlag{
			Name:  "log-format",
			Value: "text",
			Usage: "set the format used by logs ('text', or 'json')",
		},
	}
}

func runCli() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
