package config

import (
	"os"

	"github.com/ReeVicente/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitalizeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/openings.db"
	// Check if the database exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database does not exist, creating it...")
		// Create the database file and directory
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open("./db/openings.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("error connecting to SQLite: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("error migrating SQLite: %v", err)
		return nil, err
	}

	// Return the database
	return db, nil
}
