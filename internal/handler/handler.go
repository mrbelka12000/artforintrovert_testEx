// Package handler implements connections to other services and clients.
package handler

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
)

type Handler struct {
	product service.ProductStoreSrv
	l       logger.Interface
}

func NewHandler(srv service.ProductStoreSrv, log *logger.Logger) *Handler {
	return &Handler{
		product: srv,
		l:       log,
	}
}
