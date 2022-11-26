package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
)

type authHandler struct {
	authServ service.AuthService
}

func NewAuthHandler(authServ service.AuthService) *authHandler {
	return &authHandler{authServ}
}

func (a *authHandler) Login(ctx *gin.Context) {
	var body service.LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	res, err := a.authServ.Login(ctx, &body)
	if err != nil {

		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (a *authHandler) Register(ctx *gin.Context) {
	var body service.RegisterRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	res, err := a.authServ.Register(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)

}
