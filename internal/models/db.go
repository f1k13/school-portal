package db

import (
	"github.com/f1k13/school-portal/internal/logger"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
)

var DB *sqlx.DB

func ConnectDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	logger.Log.Info(dbURL, "DB URL")
	db, err := sqlx.Connect("", dbURL)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"database_url": dbURL}).Error("Fail to connection database")
		return err
	}
	DB = db
	logger.Log.Info("Success connect to database")
	return nil
}
