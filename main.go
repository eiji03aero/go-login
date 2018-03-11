package main

import (
	"os"

	"github.com/eiji03aero/go-login/action"
	"github.com/eiji03aero/go-login/controller"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-login"
	app.Usage = "demo version of login feature"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "Global option dryrun",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "User related action",
			Action:  controller.UserController,
		},
	}

	app.Before = action.AppBefore

	app.After = action.AppAfter

	app.Run(os.Args)
}
