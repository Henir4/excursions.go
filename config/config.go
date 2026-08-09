package config

import (
	"fmt"

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