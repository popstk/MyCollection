package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(os.Stdout)
	app := cli.NewApp()
	app.Name = "My Collection"
	app.Usage = "Jav"
	app.Version = "0.01"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "import",
			Usage: "import list from file",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "star,s",
					Value: 0,
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("import file: ", c.Args().First())
				return ImportFromFiles(c.Args().First(), c.Int("star"))
			},
		},
		cli.Command{
			Name:  "add",
			Usage: "add key",
			Action: func(c *cli.Context) error {
				return AddName(c.Args(), c.Int("star"))
			},
		},
		cli.Command{
			Name:  "search",
			Usage: "search key value",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "all,a",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Bool("all") {
					ShowAll(c.Int("star"))
					return nil
				}
				if c.NArg() > 0 {
					Search(c.Args())
					return nil
				}
				return errors.New("Need Args or Flags")
			},
		},
		cli.Command{
			Name:  "format",
			Usage: "format names",
			Action: func(c *cli.Context) error {
				for _, raw := range c.Args() {
					name, err := FormatName(raw)
					if err != nil {
						fmt.Println(raw, ": ", err)
						continue
					}
					fmt.Println(name)
				}
				return nil
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "database, d",
			Usage: "set database file `path`",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
