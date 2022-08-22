package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/handler"
)

func SetUpMux(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/product", h.GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/delete/{id}", h.RemoveProduct).Methods(http.MethodDelete)
	r.HandleFunc("/product/update", h.UpdateProduct).Methods(http.MethodPut)
	return r
}
