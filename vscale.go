package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

const (
	ver = "0.1.0"
)

var (
	token = ""
)

func main() {
	if os.Getenv("VSCALE") != "" {
		token = os.Getenv("VSCALE")
	} else {
		fmt.Println("VSCALE environment variable is not set")
	}
	app := cli.NewApp()
	app.Name = "vscale"
	app.Usage = "vscale.io command line interface"
	app.Version = ver
	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "account",
			Aliases: []string{"a"},
			Usage:   "Gets account info",
			Action: func(c *cli.Context) error {
				getAccountInfo()
				return nil
			},
		},
		{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "Lists ssh keys",
			Action: func(c *cli.Context) error {
				listKeys()
				return nil
			},
		},
		{
			Name:    "scalet",
			Aliases: []string{"s"},
			Usage:   "Manages scalets",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "Gets the list of scalets",
					Action: func(c *cli.Context) error {
						listScalet()
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "Creates a new scalet",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "rplan", Value: "small", Usage: "set the rplan"},
						cli.BoolFlag{Name: "do_start", Usage: "set the do_start"},
						cli.StringFlag{Name: "location", Value: "msk0", Usage: "set the location"},
						cli.StringFlag{Name: "key", Value: "test", Usage: "set the SSHKey"},
					},
					Action: func(c *cli.Context) error {
						/*fmt.Println("rplan:", c.String("rplan"))
						fmt.Println("do_start:", c.String("do_start"))
						fmt.Println("location:", c.String("location"))
						fmt.Println("key:", c.String("key"))*/
						createScalet()
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "Removes a scalet",
					Flags: []cli.Flag{
						cli.Int64Flag{Name: "id", Usage: "the ID of scalet to remove"},
					},
					Action: func(c *cli.Context) error {
						removeScalet(c.Int64("id"))
						return nil
					},
				},
				{
					Name:  "restart",
					Usage: "Restarts a scalet",
					Flags: []cli.Flag{
						cli.Int64Flag{Name: "id", Usage: "the ID of scalet to restart"},
					},
					Action: func(c *cli.Context) error {
						restartScalet(c.Int64("id"))
						return nil
					},
				},
				{
					Name:  "stop",
					Usage: "Stops a scalet",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "id", Usage: "the ID of scalet to list to stop"},
					},
					Action: func(c *cli.Context) error {
						stopScalet(c.Int64("id"))
						return nil
					},
				},
				{
					Name:  "start",
					Usage: "Stops a scalet",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "id", Usage: "the ID of scalet to start"},
					},
					Action: func(c *cli.Context) error {
						startScalet(c.Int64("id"))
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
