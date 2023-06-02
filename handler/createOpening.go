package handler

import (
	"net/http"

	"github.com/ReeVicente/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	logger.Infof("request received: %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating opening on database",
		})
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
