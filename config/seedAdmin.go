package config

import (
	"os"

	"excursion.com/schemas"
	"gorm.io/gorm"
)

// SeedAdmin checks if an admin user exists in the database, and if not, creates a default admin user with predefined credentials.
func SeedAdmin(db *gorm.DB) {
	var count int64

	adminID, err := PrefixID("adm")
	if err != nil {
		logger.Errorf("Failed to generate admin ID: %v", err)
		return
	}

	db.Model(schemas.User{}).Where("isAdmin = ?", true).Count(&count)
	if count == 0 {
		// Create default admin user
		admin := schemas.User{
			Username: os.Getenv("USERNAME"),
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			UserID:   adminID,
			IsAdmin:  true,
		}
		
		// Save the admin user to the database
		if err := db.Create(&admin).Error; err != nil {
			logger.Errorf("Failed to seed admin user: %v", err)
			return
		}
	}
}