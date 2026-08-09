package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

// ShowUsersHandler retrieves and returns a list of all users from the database.
func ShowUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	// Query the database to retrieve all users
	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	// Return the list of users in the response
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}