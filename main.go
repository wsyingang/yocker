package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.App{}
	app.Name = "hello world"
	app.Usage = "usage"
	var m_port int
	// add flag
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port,p",
			Value:       8080,
			Usage:       "listen port",
			Destination: &m_port,
		},
	}
	// Action is the default command for app
	//
	app.Action = func(c *cli.Context) {
		fmt.Println("func")
		port := c.Int("port")
		fmt.Printf("port=%d\n", port)
	}
	// Command use
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "calc 1+1",
			Action: func(c *cli.Context) error {
				fmt.Println("add func")
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "addS",
					Usage: "sub command for add",
					Action: func(c *cli.Context) {
						fmt.Println("sub command")
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
