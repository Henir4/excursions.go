package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

// UpdateExcursionHandler handles requests to update an existing excursion.
func UpdateExcursionHandler(ctx *gin.Context) {
  req := UpdateExcursionRequest{}

  ctx.BindJSON(&req)

  // Validate the request to ensure at least one field is provided for the update.
  if err := req.Validate(); err != nil {
    logger.Errorf("validation erro: %v", err.Error())
    sendError(ctx, http.StatusBadRequest, err.Error())
    return
  }

  // Retrieve the excursion ID from the query parameters and check if it's provided.
  id :=   ctx.Query("id")
  if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

  excursion := schemas.Excursion{}

  // Query the database to retrieve the excursion with the specified ID. If not found, return an error.
  if err := db.First(&excursion, id).Error; err != nil {
    sendError(ctx, http.StatusInternalServerError, "excursion not found")
    return
  }

  // Update the excursion fields with the provided values from the request, if they are not empty.
  if req.Image != "" {
    excursion.Image = req.Image
  }
  if req.Title != "" {
    excursion.Title = req.Title
  }
  if req.Description != "" {
    excursion.Description = req.Description
  }
  if req.Buy != "" {
    excursion.Buy = req.Buy
  }
  if req.FindMore != "" {
    excursion.FindMore = req.FindMore
  }

  // Save the updated excursion to the database. If there's an error, log it and return an error response.
  if err := db.Save(&excursion).Error; err != nil {
    logger.Errorf("error updating excursion %v", err.Error())
    sendError(ctx, http.StatusInternalServerError, "error updating excursion")
    return
  }

  // Return a success response with the updated excursion details.
  ctx.JSON(http.StatusOK, gin.H {
    "updated data": excursion,
  })
}