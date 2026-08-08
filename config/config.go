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
