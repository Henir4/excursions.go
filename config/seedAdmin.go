package config

import (
	"os"

	"excursion.com/schemas"
	"gorm.io/gorm"
)

func getAdminEnv(primary, fallback string) string {
	value := os.Getenv(primary)
	if value == "" {
		value = os.Getenv(fallback)
	}
	return value
}

// SeedAdmin checks if an admin user exists in the database, and if not, creates a default admin user with predefined credentials.
func SeedAdmin(db *gorm.DB) {
	var count int64

	adminID, err := PrefixID("adm")
	if err != nil {
		logger.Errorf("Failed to generate admin ID: %v", err)
		return
	}

	db.Model(schemas.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		// Create default admin user
		admin := schemas.User{
			Username: getAdminEnv("ADMIN_USERNAME", "USERNAME"),
			UserID:   adminID,
			Password: getAdminEnv("ADMIN_PASSWORD", "PASSWORD"),
			Email:    getAdminEnv("ADMIN_EMAIL", "EMAIL"),
			Role:     "admin",
		}

		if admin.Username == "" || admin.Password == "" || admin.Email == "" {
			logger.Errorf("Admin credentials missing or incomplete in environment variables")
			return
		}

		// Save the admin user to the database
		if err := db.Create(&admin).Error; err != nil {
			logger.Errorf("Failed to seed admin user: %v", err)
			return
		}

		logger.Info("Default admin user created successfully.")
	} else {
		logger.Info("Admin user already exists. Skipping seeding.")
	}
}
