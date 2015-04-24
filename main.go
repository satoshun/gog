package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gog"
	app.Version = "0.3.0"
	app.Usage = "structure directory like Go"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "base, b",
			Usage: "define git path",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "get",
			ShortName: "g",
			Usage:     "clone repository",
			Action:    actionGet,
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "update repository",
			Action:    actionUpdate,
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "list cloned repository",
			Action:    actionList,
		},
	}

	app.Run(os.Args)
}
