package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gobwas/glob"
	"github.com/pkg/errors"
	"github.com/tidwall/buntdb"
)

var dbFile string

func init() {
	dbFile = os.Getenv("COLLECTION_DB")
	if dbFile == "" {
		panic("env $COLLECTION_DB is empty")
	}
}

// Item sth items
type Item struct {
	ID   string `json:"id"`
	Star int    `json:"star"`
}


func delFromDB(db *buntdb.DB, raw string) error {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	name, err := FormatName(raw)
	if err != nil {
		fmt.Println(raw, ": ", err)
		return nil
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(name)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func updateToDB(db *buntdb.DB, raw string, star int) error {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	name, err := FormatName(raw)
	if err != nil {
		fmt.Println(raw, ": ", err)
		return nil
	}
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
		_, _, err = tx.Set(name, string(bytes[:]), nil)
		if err != nil {
			fmt.Println("Can not insert: ", name)
			return err
		}
		str, err := PrettyJSON(bytes)
		if err == nil {
			fmt.Println(str)
		}
		return nil
	})
	return nil
}

// ImportFromFiles import id from files
func ImportFromFiles(path string, star int) error {
	db, err := getSession()
	if err != nil {
		fmt.Println("Error open: ", err)
		return nil
	}
	defer db.Close()

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error open: ", err)
		return err
	}

	buf := bufio.NewReader(f)
	for {
		bytes, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		_ = updateToDB(db, string(bytes[:]), star)
	}

	return nil
}

// AddName add name single
func AddName(names []string, star int) error {
	db, err := getSession()
	if err != nil {
		fmt.Println("Error open db : ", err)
		return err
	}
	defer db.Close()
	for _, name := range names {
		_ = updateToDB(db, name, star)
	}
	return nil
}

// DelName -
func DelName(names []string) error {
	db, err := getSession()
	if err != nil {
		fmt.Println("Error open db : ", err)
		return err
	}
	defer db.Close()
	for _, name := range names {
		_ = delFromDB(db, name)
	}
	return nil
}

// Search id
func Search(star int, words []string) error {
	db, err := getSession()
	if err != nil {
		fmt.Println("Error open: ", err)
		return nil
	}
	defer db.Close()

	fmt.Println("star is ", star)

	for _, key := range words {
		fmt.Println(">>> ", key)
		name := ValidName(key)
		g := glob.MustCompile(name)
		var pivot string
		if star >= 0 {
			pivot = fmt.Sprintf(`{"star": %d}`, star)
		}
		_ = db.View(func(tx *buntdb.Tx) error {
			err = tx.AscendEqual("star", pivot, func(key, value string) bool {
				if g.Match(key) {
					value, err = PrettyJSON([]byte(value))
					if err != nil {
						fmt.Println(name, ": ", err)
					}
					fmt.Println(value)
				}
				return true
			})
			return err
		})
	}
	return nil
}

func getSession() (*buntdb.DB, error) {
	db, err := buntdb.Open(dbFile)
	if err != nil {
		return nil, errors.Wrap(err, "open buntdb")
	}

	err = db.CreateIndex("star", "*", buntdb.IndexJSON("star"))
	if err != nil {
		return nil, err
	}
	err = db.CreateIndex("actress", "*", buntdb.IndexJSON("actress"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

