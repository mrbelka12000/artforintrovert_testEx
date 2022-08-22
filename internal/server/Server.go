package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
)

const timeout = 15 * time.Second

func NewServer(r *mux.Router) *http.Server {
	cfg := config.GetConf()

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Api.Port
	}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
		WriteTimeout: timeout,
	}
}
