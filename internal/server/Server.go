package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
)

const timeout = 15 * time.Second

func NewServer(r *mux.Router) *http.Server {
	port := tools.GetApiPort()

	return &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
		WriteTimeout: timeout,
	}
}
