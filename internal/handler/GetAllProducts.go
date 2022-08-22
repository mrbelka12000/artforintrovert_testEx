package handler

import (
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := h.srv.GetAll()
	if err != nil {
		zap.S().Errorf("failed to get all products: %v", err)
		WriteResponse(w, http.StatusInternalServerError, "no data")
		return
	}

	w.Write([]byte(tools.GetJsonString(data)))
}
