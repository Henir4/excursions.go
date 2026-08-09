package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

func ShowExcursionHandler(ctx *gin.Context) {
  id := ctx.Query("id")

  excursion := schemas.Excursion{}

  // Query the database to retrieve the excursion with the specified ID
  if err := db.First(&excursion, id).Error; err != nil {
    sendError(ctx, http.StatusNotFound, "excursion not found!")
    return
  }

  // Return the excursion details in the response
  ctx.JSON(http.StatusOK, gin.H {
    "data": excursion,
  })
}