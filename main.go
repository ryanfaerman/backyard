package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ryanfaerman/picket/command"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	app := cli.NewApp()
	app.Name = "picket"

	app.Commands = []cli.Command{
		command.Blacklist,
	}

	app.Run(os.Args)
	return 0
}
