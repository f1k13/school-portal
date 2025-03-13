package main

import (
	"github.com/f1k13/school-portal/internal/app"
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("Error loading .env files")
	}
	app.App()
}
