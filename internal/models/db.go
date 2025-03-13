package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		return err
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Не удалось пинговать базу данных: %v", err)
		return err
	}

	log.Println("Успешно подключено к базе данных")
	return nil
}
