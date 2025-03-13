package app

import (
	"net/http"

	"github.com/f1k13/school-portal/internal/logger"
	db "github.com/f1k13/school-portal/internal/models"
	"github.com/f1k13/school-portal/internal/routes"
	"github.com/go-chi/chi/v5"
)

func App() {
	initConnectDB()
	initRouter()
}
func initConnectDB() {
	if err := db.ConnectDB(); err != nil {
		logger.Log.Error("Failed connection to db", err)
	}
}
func initRouter() {
	r := chi.NewRouter()
	routes.StartRouter(r)
	logger.Log.Info("SERVER START ON PORT", 3000)
	http.ListenAndServe(":3000", r)
}
