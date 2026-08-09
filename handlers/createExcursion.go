package handler

import (
	"net/http"

	"excursion.com/config"
	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

// Validate Function for Request
func CreateExcursionHandler(ctx *gin.Context) {
	var req CreateExcursionRequest

	// Attempt to bind the JSON payload to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "malformed JSON: " + err.Error()})
		return
	}

	// Validate the decoded struct
	if err := req.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	excursionID, err := config.PrefixID("exc")
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error generating excursion ID")
		return
	}


	excursion := schemas.Excursion{
		ExcursionID: excursionID,
		Image:       req.Image,
		Title:       req.Title,
		Description: req.Description,
		Buy:         req.Buy,
		FindMore:    req.FindMore,
	}

	
	// Insert data in the database
	if err := db.Create(&excursion).Error; err != nil {
		logger.Errorf("error creating excursion: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating excursion on database")
		return
	}

	logger.Infof("request received: %+v", req)
	// If everything is valid, proceed
	sendSuccess(ctx, "excursion-created", excursion)
}
