package config

import (
	"os"

	"excursion.com/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func IntializeSQLite() (*gorm.DB, error) {
  // Logs
  logger := GetLogger("sqlite")
  // Check if Database Exists
  dbPath := "./db/excursions.db" 
  _, err := os.Stat(dbPath)
  // If database file doesn't exist, create a new one
  if os.IsNotExist(err){

    logger.Info("database Not Found")
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

  // Create and Connect to Database
  db, err := gorm.Open(sqlite.Open("./db/excursions.db"),  &gorm.Config{})
  // Connection Error Log
  if err != nil {
  logger.Errorf("sqlite initialization error: %v", err)
    return nil, err
  }

  // Migrate the database schema for Excursion and User models
  err = db.AutoMigrate(&schemas.Excursion{}, &schemas.User{})
  if err != nil {
    // Migration Error Log
    logger.Errorf("sqlite migration error: %v", err)
    return nil, err
  }

    return db, err
}
