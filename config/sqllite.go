package config

import (
	"ms-gin-gonic-rest-api-go/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if db already exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database not exists, creating new one...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// DB Connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite connection error: %v", err)
		return nil, err
	}

	// DB Migration
	err = db.AutoMigrate(&schemas.Openning{})

	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
	}

	return db, nil
}
