package handler

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
)

type Handler struct {
	product service.ProductStoreSrv
}

func NewHandler(srv service.ProductStoreSrv) *Handler {
	return &Handler{product: srv}
}
