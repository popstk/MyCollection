package main

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

var prefixs []string

func init() {
	f, err := os.Open("format.json")
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = json.NewDecoder(f).Decode(&v)
	if err != nil {
		panic(err)
	}

	root, ok := v.(map[string]interface{})
	if !ok {
		panic("Invalid Format: root")
	}

	prefix, ok := root["prefix"].([]interface{})
	if !ok {
		panic("Invalid Format: prefix")
	}

	for _, p := range prefix {
		pre, ok := p.(string)
		if !ok {
			panic("Invalid Format: prefix value")
		}
		prefixs = append(prefixs, pre)
	}
}

// FormatName format names
func FormatName(name string) (string, error) {
	name = strings.ToUpper(name)

	for _, sep := range []byte{'-', '_', ' '} {
		name = strings.Replace(name, string(sep), "", -1)
	}

	prefix := ""
	for _, pre := range prefixs {
		if strings.HasPrefix(name, strings.ToUpper(pre)) {
			prefix = pre
			break
		}
	}

	if prefix == "" {
		return "", errors.New("no match prefix")
	}

	rest := name[len(prefix):]
	n, err := strconv.Atoi(rest)
	if err != nil {
		return "", errors.New("not number:" + rest)
	}

	return prefix + "-" + strconv.Itoa(n), nil
}
