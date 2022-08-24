// Package tools contains helper functions.
package tools

import (
	"bytes"
	"encoding/json"
)

func GetJsonString(value interface{}) string {
	if value == nil {
		return "{}"
	}
	bf := bytes.NewBufferString("")
	e := json.NewEncoder(bf)
	e.SetEscapeHTML(false)
	e.Encode(value)
	return bf.String()
}
