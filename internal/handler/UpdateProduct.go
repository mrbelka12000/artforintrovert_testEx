package handler

import (
	"encoding/json"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var product *models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		zap.S().Errorf("invalid body: %v", err)
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = product.Validate()
	if err != nil {
		zap.S().Errorf("invalid product settings: %v", err)
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.srv.Update(product)
	if err != nil {
		status, msg := parseErrorResponse(err)
		WriteResponse(w, status, msg)
		return
	}

	WriteResponse(w, http.StatusOK, "updated")
}
