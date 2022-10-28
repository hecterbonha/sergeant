package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hecterbonha/sergeant/db"
	"github.com/hecterbonha/sergeant/lib/shared"
	"github.com/hecterbonha/sergeant/model"
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

type GenerateHandlerResponse struct {
	Message string       `json:"message"`
	Payload []model.User `json:"payload"`
	Status  string       `json:"status"`
}

func Generate(ctx *gin.Context) {
	users := []model.User{}
	db.Postgres.Find(&users)
	ctx.JSON(http.StatusOK, &GenerateHandlerResponse{
		Message: "Handshake is good and everything",
		Payload: users,
		Status:  shared.SuccessStatusCode,
	})
}
