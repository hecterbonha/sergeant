package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/lib/shared"
)

type HandshakePayload struct {
	Token string `json:"token"`
}

type HandshakeHandlerResponse struct {
	Message string           `json:"message"`
	Payload HandshakePayload `json:"payload"`
	Status  string           `json:"status"`
}

func HandShake(ctx *gin.Context) {
	token := GetPayloadToken()

	ctx.JSON(http.StatusOK, &HandshakeHandlerResponse{
		Message: "Handshake is good and everything",
		Payload: HandshakePayload{
			Token: token,
		},
		Status: shared.SuccessStatusCode,
	})
}
