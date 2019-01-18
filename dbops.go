package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tidwall/buntdb"
)

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
		updateToDB(db, string(bytes[:]), star)
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
		updateToDB(db, name, star)
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
		delFromDB(db, name)
	}
	return nil
}

// Search id
func Search(words []string) error {
	db, err := getSession()
	if err != nil {
		fmt.Println("Error open: ", err)
		return nil
	}
	defer db.Close()
	for _, key := range words {
		name, err := FormatName(key)
		if err != nil {
			fmt.Println(name, ": ", err)
			continue
		}
		db.View(func(tx *buntdb.Tx) error {
			v, err := tx.Get(name)
			if err != nil {
				fmt.Println(name, ": ", err)
				return err
			}

			v, err = PrettyJSON([]byte(v))
			if err != nil {
				fmt.Println(name, ": ", err)
			}
			fmt.Println(v)
			return nil
		})
	}
	return nil
}

// ShowAll show all names
func ShowAll(star int) error {
	db, err := getSession()
	if err != nil {
		fmt.Println("Error open db : ", err)
		return err
	}
	defer db.Close()

	db.View(func(tx *buntdb.Tx) error {
		pivot := fmt.Sprintf(`{"star": %d}`, star)
		tx.AscendEqual("star", pivot, func(key, value string) bool {
			v, err := PrettyJSON([]byte(value))
			if err != nil {
				fmt.Println(key, ": ", err)
			}
			fmt.Println(v)
			return true
		})
		return nil
	})

	return nil
}

func getSession() (*buntdb.DB, error) {
	db, err := buntdb.Open("jav.db")
	if err != nil {
		return nil, err
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

// ExportToFile export id to file
func ExportToFile(path string, star int) error {
	return nil
}
