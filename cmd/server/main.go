package main

import (
	"github.com/f1k13/school-portal/internal/logger"
	db "github.com/f1k13/school-portal/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"sync"
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
	r := gin.Default()
	logger.Log.Info("SERVER START ON PORT", 3000)
	if err := r.Run(`:3000`); err != nil {
		logger.Log.Fatal("Error starting server", err)
	}
}
