package main

import (
	"bytes"
	"encoding/json"
)

// PrettyJson format json string
func PrettyJson(str []byte) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, str, "", " ")
	return out.String(), err
}
