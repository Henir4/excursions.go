package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

func ShowExcursionsHandler(ctx *gin.Context) {
  excursions := []schemas.Excursion{}

  // Query the database to retrieve all excursions
  if err := db.Find(&excursions).Error; err != nil {
    sendError(ctx, http.StatusInternalServerError, "error listing excursions")
    return
  }

  // Return the list of excursions in the response
  ctx.JSON(http.StatusOK, gin.H {
    "data": excursions,
  })
}