package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/secrets-bridge-v2/server"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func beforeApp(c *cli.Context) error {
	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "secrets-bridge-v2"
	app.Version = VERSION
	app.Usage = "Rancher Vault Volume Driver"
	app.Before = beforeApp
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "debug,d",
		},
	}

	app.Commands = []cli.Command{
		server.Command(),
	}

	app.Run(os.Args)
}
