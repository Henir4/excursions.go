package handler

import (

	"github.com/gin-gonic/gin"
)

func AdminDashboardHandler(ctx *gin.Context) {
	// Placeholder for admin dashboard logic
	// This function can be expanded to include specific admin functionalities.
	sendSuccess(ctx, "admin-dashboard", gin.H{"message": "Welcome to the Admin Dashboard!"})
}