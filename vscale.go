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

func main() {
	app := cli.NewApp()
	app.Name = "vscale"
	app.Usage = "vscale.io command line interface"
	app.Version = ver
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
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
			Aliases: []string{"c"},
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
						scaletList()
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "Creates a new scalet",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "Removes a scalet",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "get",
					Usage: "Gets a scalet",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "restart",
					Usage: "Restarts a scalet",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "stop",
					Usage: "Stops a scalet",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "start",
					Usage: "Stops a scalet",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
