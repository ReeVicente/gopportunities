package handler

import (
	"fmt"
	"net/http"

	"github.com/ReeVicente/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Update Opening
	if err := db.Model(&opening).Updates(request).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opening width id: %s", id))
		return
	}

	sendSuccess(ctx, "update-opening", opening)

}
