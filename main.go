package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sdomino/scribble"
	"github.com/tidwall/buntdb"
	"github.com/urfave/cli"
)

func testbuntdb() {
	db, err := buntdb.Open("jav.db")
	if err != nil {
		fmt.Println("Error open: ", err)
		return
	}
	defer db.close()
	db.CreateIndex("last_name", "*", buntdb.IndexJSON("name.last"))
}

func getconnection() (*Driver, error) {
	db, err := scribble.New("db", nil)
	if err != nil {
		fmt.Println("Error", err)
		return nil, err
	}

	return db, nil
}

func main() {
	log.SetOutput(os.Stdout)
	app := cli.NewApp()
	app.Name = "My Collection"
	app.Usage = "Jav"
	app.Version = "0.01"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "import",
			Usage: "import list",
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}
	app.Run(os.Args)
}
