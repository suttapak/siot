package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
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
