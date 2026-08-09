package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

// sendError sends a JSON error response with the specified HTTP status code and message.
func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H {
		"message": msg,
		"errorCode": code,
	})
}

// sendSuccess sends a JSON success response with the specified message and data.
func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H {
		"message": op,
		"data": data,
	})
}

// Response structures for various API responses
type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode int `json:"errorCode"`
}

// CreateExcursionResponse represents the response structure for creating an excursion.
type CreateExcursionResponse struct {
	Message string `json:"message"`
	Data schemas.ExcursionResponse `json:"data"`
}

// DeleteExcursionResponse represents the response structure for deleting an excursion.
type DeleteExcursionResponse struct {
	Message string `json:"message"`
	Data schemas.ExcursionResponse `json:"data"`
}

// ShowExcursionResponse represents the response structure for showing a single excursion.
type ShowExcursionResponse struct {
	Message string `json:"message"`
	Data schemas.ExcursionResponse `json:"data"`
}

// ShowExcursionsResponse represents the response structure for showing multiple excursions.
type ShowExcursionsResponse struct {
	Message string `json:"message"`
	Data []schemas.ExcursionResponse `json:"data"`
}

// UpdateExcursionResponse represents the response structure for updating an excursion.
type UpdateExcursionResponse struct {
	Message string `json:"message"`
	Data schemas.ExcursionResponse `json:"data"`
}