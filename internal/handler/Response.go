package handler

import (
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
	"net/http"
)

type errorResp struct {
	Status int
	Error  string
}

func ErrorResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Write([]byte(tools.GetJsonString(&errorResp{Status: code, Error: msg})))
}
