package tools

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
)

func GetApiPort() (port string) {
	cfg := config.GetConf()
	port = os.Getenv("PORT")
	if port == "" {
		port = cfg.Api.Port
	}
	return
}

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
