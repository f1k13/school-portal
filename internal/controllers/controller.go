package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/f1k13/school-portal/internal/middleware/auth"
)

type Response struct {
	Message string `json:"message"`
}
type Controller struct{}

func (h *Controller) ResponseJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error in method ResponseJson", http.StatusBadRequest)
	}
}

func (h *Controller) GetUserIDCtx(ctx context.Context) string {
	userID, _ := ctx.Value(auth.UserIDKey).(string)
	return userID
}
