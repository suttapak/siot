package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
)

type adminUserHandler struct {
	userServ service.UserService
}

func NewAdminUserHandler(userServ service.UserService) *adminUserHandler {
	return &adminUserHandler{userServ}
}

func (h *adminUserHandler) Users(ctx *gin.Context) {
	res, err := h.userServ.FindUsers(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *adminUserHandler) AddRoles(ctx *gin.Context) {
	body := service.AddRolesRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	res, err := h.userServ.AddRoles(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
