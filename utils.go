package main

import (
	"bytes"
	"encoding/json"
)

func PrettyJson(str []byte) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, str, "", " ")
	return out.String(), err
}
