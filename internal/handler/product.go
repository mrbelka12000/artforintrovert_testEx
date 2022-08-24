package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/models"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/tools"
)

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := h.product.GetAllProducts()
	if err != nil {
		h.l.Errorf("failed to get all products: %v", err)
		WriteResponse(w, http.StatusInternalServerError, "no data")
		return
	}

	w.Write([]byte(tools.GetJsonString(data)))
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		h.l.Warn("no id to delete")
		WriteResponse(w, http.StatusBadRequest, "no id to delete")
		return
	}

	err := h.product.DeleteProduct(r.Context(), id)
	if err != nil {
		status, msg := service.ParseErrorResponse(err)
		WriteResponse(w, status, msg)
		return
	}

	h.l.Infof("document %v successfully deleted", id)
	WriteResponse(w, http.StatusNoContent, "deleted")
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		h.l.Errorf("invalid body: %v", err)
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = product.Validate()
	if err != nil {
		h.l.Errorf("invalid product settings: %v", err)
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.product.UpdateProduct(r.Context(), product)
	if err != nil {
		status, msg := service.ParseErrorResponse(err)
		WriteResponse(w, status, msg)
		return
	}

	h.l.Infof("document %v successfully updated", product.ID.String())
	WriteResponse(w, http.StatusOK, "updated")
}
