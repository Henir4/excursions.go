package handler

import "github.com/gin-gonic/gin"

// sendError sends a JSON error response with the specified HTTP status code and message.
func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H {
		"message": msg,
		"errorCode": code,
	})
}