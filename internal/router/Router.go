package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/handler"
)

func SetUpMux(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/remove/{id}", h.RemoveProduct).Methods(http.MethodDelete)
	r.HandleFunc("/update", h.UpdateProduct).Methods(http.MethodPost)
	return r
}
