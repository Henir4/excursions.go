package handler

import (
	"fmt"
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

// DeleteExcursionHandler handles requests to delete an existing excursion.
func DeleteExcursionHandler(ctx *gin.Context) {
  id := ctx.Query("id")
  if id == "" {
    sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
    return
  }
  
  excursion := schemas.Excursion{}

  // Query the database to retrieve the excursion with the specified ID. If not found, return an error.
  if err := db.First(&excursion, id).Error; err != nil {
    sendError(ctx, http.StatusNotFound, fmt.Sprintf("excursion with id: %s", id))
    return
  }

  // Delete the excursion from the database. If there's an error, log it and return an error response.
  if err := db.Delete(&excursion).Error; err != nil {
    sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro deleting excursion with id: %s", id))
    return
  }


  // Return a success response with the deleted excursion details.
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "excursion deleted successfully",
    "data": excursion,
  })
}
