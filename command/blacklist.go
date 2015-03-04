package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ryanfaerman/picket/picket"
)

var blacklist map[string][]string

var Blacklist = cli.Command{
	Name:  "blacklist",
	Usage: "Manage the blacklist",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "remote, r",
			Usage: "is the url a remote source of blacklist entries",
		},
		cli.StringFlag{
			Name:  "group, g",
			Value: "base",
			Usage: "the group to add the url to",
		},
	},
	Action: func(c *cli.Context) {
		println(c.Bool("remote"))
		println(c.String("group"))

		item := c.Args().First()
		if item != "" {
			picket.AddToBlacklist(item)
		}
		fmt.Println(picket.ListBlacklist())
	},
}
