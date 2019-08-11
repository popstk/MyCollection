package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var conf Conf

// Conf config
type Conf struct {
	Delims  string   `json:"delims"`
	Ignores []string `json:"ignores"`
	Prefix  []string `json:"prefix"`
}

func init() {
	dir, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir = filepath.Dir(dir)

	f, err := os.Open(filepath.Join(dir, "format.json"))
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		panic(err)
	}
}
