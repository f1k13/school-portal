package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/f1k13/school-portal/internal/logger"
	"github.com/f1k13/school-portal/internal/routes"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func App() {
	ConnectDB()
	StartApp()
}

var DB *sql.DB

func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		logger.Log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		return err
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Не удалось пинговать базу данных: %v", err)
		return err
	}

	log.Println("Успешно подключено к базе данных")
	return nil
}

func StartApp() {

	r := chi.NewRouter()
	routes.StartRouter(r, DB)
	logger.Log.Info("SERVER START ON PORT", 3000)
	http.ListenAndServe(":3000", r)
}
