package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
)

type userHandler struct {
	userServ service.UserService
}

func NewUserHandler(userServ service.UserService) *userHandler {
	return &userHandler{userServ}
}

func (u *userHandler) FindUser(ctx *gin.Context) {
	uid, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	var body = service.FindUserRequest{
		UserId: uid,
	}
	res, err := u.userServ.FindUser(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (u *userHandler) FindUserById(ctx *gin.Context) {

	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}

	var body = service.FindUserRequest{UserId: userId}

	res, err := u.userServ.FindUser(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
