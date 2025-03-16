package handlers

import (
	"context"
	"encoding/json"
	"github.com/f1k13/school-portal/internal/middleware"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func ResponseJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error in method responseJson", http.StatusBadRequest)
	}
}

func GetUserIDCtx(ctx context.Context) string {
	userID := ctx.Value(middleware.UserIDKey).(string)
	return userID
}
