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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "database, d",
			Usage: "set database file `path`",
		},
		cli.IntFlag{
			Name:  "star,s",
			Value: -1,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "import",
			Usage: "import list from file",
			Action: func(c *cli.Context) error {
				fmt.Println("import file: ", c.Args().First())
				return ImportFromFiles(c.Args().First(), c.Int("star"))
			},
		},
		{
			Name:  "add",
			Usage: "add key",
			Action: func(c *cli.Context) error {
				return AddName(c.Args(), c.Int("star"))
			},
		},
		{
			Name:  "del",
			Usage: "del key",
			Action: func(c *cli.Context) error {
				return DelName(c.Args())
			},
		},
		{
			Name:  "find",
			Usage: "find key value",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					_ = Search(c.GlobalInt("star"), c.Args())
					return nil
				}
				return errors.New("need Args or Flags")
			},
		},
		{
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

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
