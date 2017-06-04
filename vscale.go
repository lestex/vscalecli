package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

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
						cli.StringFlag{Name: "from", Value: "ubuntu_16.04_64_001_master", Usage: "set the template for a scalet"},
						cli.StringFlag{Name: "rplan", Value: "small", Usage: "set the rplan"},
						cli.StringFlag{Name: "do_start", Value: "true", Usage: "set the do_start"},
						cli.StringFlag{Name: "location", Value: "msk0", Usage: "set the location"},
						cli.StringFlag{Name: "key", Usage: "set the SSHKey"},
						cli.StringFlag{Name: "name", Value: "", Usage: "set the hostname"},
					},
					Action: func(c *cli.Context) error {
						dostart, _ := strconv.ParseBool(c.String("do_start"))
						keys := []string{c.String("key")}
						createScalet(c.String("from"),
							c.String("rplan"),
							c.String("name"),
							c.String("location"),
							dostart,
							keys)
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
				{
					Name:    "rplans",
					Aliases: []string{"p"},
					Usage:   "Lists list rplans",
					Action: func(c *cli.Context) error {
						listRplans()
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
