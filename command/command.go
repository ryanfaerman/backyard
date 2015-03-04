package command

// import "github.com/codegangsta/cli"

// type Command interface {
// 	Usage() string
// 	Action(c *cli.Context)
// }

// type FlaggableCommand interface {
// 	Command
// 	Flags() []cli.Flag
// }

// func Factory(commands map[string]Command) (out []cli.Command) {
// 	for name, cmd := range commands {
// 		c := cli.Command{
// 			Name:   name,
// 			Usage:  cmd.Usage(),
// 			Action: cmd.Action,
// 		}

// 		if cmd.(type) == FlaggableCommand {
// 			c.Flags = cmd.(FlaggableCommand).Flags()
// 		}

// 		out = append(out, c)
// 	}

// 	return out
// }
