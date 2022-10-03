package handler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpMux(h *Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/product", h.GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/delete/{id}", h.DeleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("/product/update", h.UpdateProduct).Methods(http.MethodPut)
	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	return r
}
