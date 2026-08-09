package config

import (
	"fmt"
	"os"

	"excursion.com/schemas"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	// Database Connection
	db     *gorm.DB
	logger *Logger
)

// Databse Handler
func Init() error {
	var err error

	// Load environment variables from the .env file if present.
	if err = godotenv.Load(); err != nil {
		if logger != nil {
			logger.Warnf("Could not load .env file: %v", err)
		}
	}

	schemas.JwtKey = []byte(os.Getenv("JWT"))
	if len(schemas.JwtKey) == 0 {
		return fmt.Errorf("JWT secret missing from environment")
	}

	// Intialize Database
	db, err = IntializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %w", err)
	}

	// Run the seed function to create an admin user if it doesn't exist
	SeedAdmin(db)

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

// Logger Initializer
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
