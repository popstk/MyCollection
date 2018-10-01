package main

import (
	"bytes"
	"encoding/json"
)

// PrettyJSON format json string
func PrettyJSON(str []byte) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, str, "", " ")
	return out.String(), err
}
