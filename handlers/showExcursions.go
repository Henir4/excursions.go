package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

func ShowExcursionsHandler(ctx *gin.Context) {
  excursions := []schemas.Excursion{}

  if err := db.Find(&excursions).Error; err != nil {
    sendError(ctx, http.StatusInternalServerError, "error listing excursions")
    return
  }

  ctx.JSON(http.StatusOK, gin.H {
    "data": excursions,
  })
}