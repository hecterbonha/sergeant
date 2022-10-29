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

func PingHandler(ctx *gin.Context) {
	nameFromQuery := ctx.Query("name")
	locales := ctx.MustGet("locales").(string)
	if nameFromQuery == "" || locales == "" {
		ctx.JSON(http.StatusBadRequest, &PingHandlerResponse{
			Message: shared.MalformedRequestMessage,
			Status:  shared.ErrorStatusCode,
		})
		return
	}
	message := GetPingMessages(nameFromQuery, locales)
	ctx.JSON(http.StatusOK, &PingHandlerResponse{
		Message: message,
		Status:  shared.SuccessStatusCode,
	})
}
