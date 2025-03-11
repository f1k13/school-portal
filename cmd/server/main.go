package main

import (
	"net/http"
	"sync"

	"github.com/f1k13/school-portal/internal/logger"
	db "github.com/f1k13/school-portal/internal/models"
	"github.com/f1k13/school-portal/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("Error loading .env files")
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		initRouter()
	}()
	go func() {
		defer wg.Done()
		initConnectDB()
	}()
	wg.Wait()
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
