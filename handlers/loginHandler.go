package handler

import (
	"net/http"

	"excursion.com/middleware"
	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler handles user login requests
func LoginHandler(ctx *gin.Context) {

	var loginRequest schemas.LoginRequest

	// Bind the JSON request body to the loginRequest struct
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username and password are provided
	if loginRequest.Username == "" || loginRequest.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Find the user in the database by username
	var user schemas.User
	if err := db.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token for the authenticated user
	tokenString, err := middleware.GenerateToken(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token in the response
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})

}

// ProtectedHandler is an example of a protected route that requires authentication
func ProtectedHandler(ctx *gin.Context) {

	// Retrieve the user from the context set by the AuthMiddleware
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	// Return a success response with the authenticated user's details
	sendSuccess(ctx, "protected-access", user)
}

