package handler

import (
	"net/http"

	"excursion.com/middleware"
	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler handles user login requests
func LoginHandler(c *gin.Context) {

	var loginRequest schemas.LoginRequest

	// Bind the JSON request body to the loginRequest struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username and password are provided
	if loginRequest.Username == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Find the user in the database by username
	var user schemas.User
	if err := db.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token for the authenticated user
	tokenString, err := middleware.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

// ProtectedHandler is an example of a protected route that requires authentication
func ProtectedHandler(c *gin.Context) {

	// Retrieve the user from the context set by the AuthMiddleware
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	// Return a success message along with the user information
	c.JSON(http.StatusOK, gin.H{"message": "Access granted", "user": user})
}