package handler

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		zap.S().Warn("no id to delete")
		WriteResponse(w, http.StatusBadRequest, "no id to delete")
		return
	}

	err := h.srv.Delete(id)
	if err != nil {
		status, msg := parseErrorResponse(err)
		WriteResponse(w, status, msg)
		return
	}

	zap.S().Infof("document %v successfully deleted", id)
	WriteResponse(w, http.StatusOK, "deleted")
}
