package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tidwall/buntdb"
	"github.com/urfave/cli"
)

type Item struct {
	ID   string `json:"id"`
	Star int    `json:"star"`
}

func getSession() (*buntdb.DB, error) {
	db, err := buntdb.Open("jav.db")
	if err != nil {
		return nil, err
	}

	db.CreateIndex("star", "*", buntdb.IndexJSON("star"))
	db.CreateIndex("actress", "*", buntdb.IndexJSON("actress"))
	return db, nil
}

func testbuntdb() {
	db, err := getSession()
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
				db, err := getSession()
				if err != nil {
					fmt.Println("Error open: ", err)
					return nil
				}
				defer db.Close()
				f, err := os.Open(c.Args().First())
				if err != nil {
					fmt.Println("Error open: ", err)
					return err
				}
				fmt.Println("start")
				star := c.Int("star")
				buf := bufio.NewReader(f)
				for {
					bytes, _, err := buf.ReadLine()
					if err == io.EOF {
						break
					}
					line := string(bytes[:])
					fmt.Println("Content is: ", line)
					line = strings.TrimSpace(line)
					name, err := FormatName(line)
					if err != nil {
						fmt.Println(name, ": ", err)
						continue
					}
					fmt.Println(name)
					err = db.Update(func(tx *buntdb.Tx) error {
						data := Item{
							ID:   name,
							Star: star,
						}
						bytes, err := json.Marshal(data)
						if err != nil {
							fmt.Println("Can not Marshal: ", data)
							return err
						}
						pvalue, replaced, err := tx.Set(name, string(bytes[:]), nil)
						fmt.Println(pvalue, replaced, err)
						return nil
					})
				}

				fmt.Println("database file path is ", c.String("database"))
				return nil
			},
		},
		cli.Command{
			Name:  "search",
			Usage: "search key value",
			Action: func(c *cli.Context) error {
				db, err := getSession()
				if err != nil {
					fmt.Println("Error open: ", err)
					return nil
				}
				defer db.Close()
				for i := 0; i < c.NArg(); i++ {
					key := c.Args().Get(i)
					name, err := FormatName(key)
					if err != nil {
						fmt.Println(name, ": ", err)
						continue
					}
					db.View(func(tx *buntdb.Tx) error {
						v, err := tx.Get(name)
						if err == nil {
							fmt.Println(name, ": ", v)
						} else {
							fmt.Println(name, ": ", err)
						}
						return nil
					})
				}
				return nil
			},
		},
		cli.Command{
			Name:  "format",
			Usage: "format names",
			Action: func(c *cli.Context) error {
				db, err := getSession()
				if err != nil {
					fmt.Println("Error open: ", err)
					return nil
				}
				defer db.Close()
				for i := 0; i < c.NArg(); i++ {
					raw := c.Args().Get(i)
					name, err := FormatName(raw)
					if err != nil {
						fmt.Println(name, ": ", err)
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

	app.Action = func(c *cli.Context) error {
		fmt.Println("database file path is ", c.String("database"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
