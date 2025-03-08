package db

import (
	"fmt"
	"log"
	"os"

	"github.com/f1k13/school-portal/internal/models/user"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}
	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}
	log.Println("Successfully connected to database")
	return nil
}
 