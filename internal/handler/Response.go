package handler

import (
	"errors"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
	"go.uber.org/zap"
	"net/http"
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

func parseErrorResponse(err error) (int, string) {
	zap.S().Debugf("error: %v", err)
	if errors.Is(err, service.ErrClientError) {
		return http.StatusBadRequest, err.Error()
	} else {
		return http.StatusInternalServerError, errors.Unwrap(err).Error()
	}
}
