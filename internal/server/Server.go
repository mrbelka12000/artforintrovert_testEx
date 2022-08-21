package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
)

const timeout = 15 * time.Second

func NewServer(r *mux.Router) *http.Server {
	cfg := config.GetConf()

	return &http.Server{
		Addr:         ":" + cfg.Api.Port,
		Handler:      r,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
		WriteTimeout: timeout,
	}
}
