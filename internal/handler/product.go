package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/service"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
)

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := h.srv.Product.GetAll()
	if err != nil {
		zap.S().Errorf("failed to get all products: %v", err)
		WriteResponse(w, http.StatusInternalServerError, "no data")
		return
	}

	w.Write([]byte(tools.GetJsonString(data)))
}

func (h *Handler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		zap.S().Warn("no id to delete")
		WriteResponse(w, http.StatusBadRequest, "no id to delete")
		return
	}

	err := h.srv.Product.Delete(id)
	if err != nil {
		status, msg := service.ParseErrorResponse(err)
		WriteResponse(w, status, msg)
		return
	}

	zap.S().Infof("document %v successfully deleted", id)
	WriteResponse(w, http.StatusNoContent, "deleted")
}

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

	err = h.srv.Product.Update(product)
	if err != nil {
		status, msg := service.ParseErrorResponse(err)
		WriteResponse(w, status, msg)
		return
	}

	WriteResponse(w, http.StatusOK, "updated")
}
