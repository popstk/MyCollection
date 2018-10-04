package main

import (
	"encoding/json"
	"os"
)

var conf Conf

// Conf config
type Conf struct {
	Delims  string   `json:"delims"`
	Ignores []string `json:"ignores"`
	Prefix  []string `json:"prefix"`
}

func init() {
	f, err := os.Open("format.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		panic(err)
	}
}
