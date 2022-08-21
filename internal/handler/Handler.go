package handler

import "github.com/mrbelka12000/artforintrovert_testEx/internal/service"

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{srv}
}
