package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/services"
)

type Handler struct {
	Service *services.Service
}

func NewHandler() *Handler {
	return &Handler{}
}
func ResponseJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error in method responseJson", http.StatusBadRequest)
	}
}
