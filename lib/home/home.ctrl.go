package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/lib/shared"
)

type PingHandlerResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type PingHandlerErrorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func PingHandler(ctx *gin.Context) {
	nameFromQuery := ctx.Query("name")
	locales := ctx.MustGet("locales").(string)

	if nameFromQuery == "" {
		ctx.JSON(http.StatusBadRequest, &PingHandlerErrorResponse{
			Message: "MALFORMED_REQUEST_FORMAT",
			Status:  shared.ErrorCode,
		})
		return
	}

	message := GetPingMessages(nameFromQuery, locales)

	ctx.JSON(http.StatusOK, &PingHandlerResponse{
		Message: message,
		Status:  shared.SuccessCode,
	})
}
