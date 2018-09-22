package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tidwall/buntdb"
	"github.com/urfave/cli"
)

func testbuntdb() {
	db, err := buntdb.Open("jav.db")
	if err != nil {
		fmt.Println("Error open: ", err)
		return
	}
	defer db.Close()
	err = db.Update(func(tx *buntdb.Tx) error {
		pvalue, replaced, err := tx.Set("iptd799", `{"id":"iptd799", "star":"5", "acc"}`, nil)
		fmt.Println(pvalue, replaced, err)
		return nil
	})
	db.View(func(tx *buntdb.Tx) error {
		v, err := tx.Get("iptd799")
		fmt.Println(v, err)
		return nil
	})
}

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
				fmt.Println("database file path is ", c.String("database"))
				return nil
			},
		},
		cli.Command{
			Name:  "search",
			Usage: "search key value",
			Action: func(c *cli.Context) error {
				fmt.Println("search file: ", c.Args().First())
				return nil
			},
		},
		cli.Command{
			Name:  "format",
			Usage: "format names",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.NArg(); i++ {
					raw := c.Args().Get(i)
					fmt.Println("Before is ", raw)
					fmt.Println("After is ", FormatName(raw))
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

	app.Action = func(c *cli.Context) error {
		fmt.Println("fuck")
		fmt.Println("database file path is ", c.String("database"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
