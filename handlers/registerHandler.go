package handler

import (
	"errors"
	"net/http"

	"excursion.com/config"
	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterHandler handles user registration requests
func RegisterHandler(ctx *gin.Context) {
	var registerRequest schemas.RegisterRequest

	// Bind the JSON request body to the registerRequest struct
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username, email, and password are provided
	if registerRequest.Username == "" || registerRequest.Email == "" || registerRequest.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username, email, and password are required"})
		return
	}

	// Check if the username or email already exists in the database
	var existing schemas.User
	err := db.Where("username = ? OR email = ?", registerRequest.Username, registerRequest.Email).First(&existing).Error
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
		return
	}

	// If the error is not a "record not found" error, return an internal server error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user"})
		return
	}

	// Generate a unique user ID using the PrefixID function from the config package
	userID, err := config.PrefixID("usr")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate user ID"})
		return
	}

	// Create a new user instance with the provided details and the generated user ID
	user := schemas.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		UserID:   userID,
		IsAdmin:  false,
	}

	// Save the new user to the database
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return a success response with the created user details
	sendSuccess(ctx, "user-registered", user)
}
