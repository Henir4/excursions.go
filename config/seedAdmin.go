package config

import (
	"excursion.com/schemas"
	"gorm.io/gorm"
)

// SeedAdmin checks if an admin user exists in the database, and if not, creates a default admin user with predefined credentials.
func SeedAdmin(db *gorm.DB) {
	var count int64

	db.Model(schemas.User{}).Where("isAdmin = ?", true).Count(&count)
	if count == 0 {
		// Create default admin user
		admin := schemas.User{
			Email:    "admin@example.com",
			Password: "admin123",
			IsAdmin:  true,
		}
		
		// Save the admin user to the database
		if err := db.Create(&admin).Error; err != nil {
			logger.Errorf("Failed to seed admin user: %v", err)
			return
		}
	}
}