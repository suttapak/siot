package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/errs"
)

type canSubHandler struct {
	canSubServ service.CanSubService
}

func NewCanSubHandler(canSubServ service.CanSubService) *canSubHandler {
	return &canSubHandler{canSubServ}
}

func (h *canSubHandler) CanSub(ctx *gin.Context) {
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(500, errs.ErrBadRequest)
		return
	}
	boxSecret := ctx.Param("boxSecret")
	res, err := h.canSubServ.CanSub(ctx, boxId, boxSecret)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
