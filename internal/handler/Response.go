package handler

import (
	"net/http"

	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
)

type Resp struct {
	Status int    `json:"status"`
	Text   string `json:"text"`
}

func WriteResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(tools.GetJsonString(&Resp{Status: code, Text: msg})))
}
