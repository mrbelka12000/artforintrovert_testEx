package handler

import "net/http"

func (h *Handler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
