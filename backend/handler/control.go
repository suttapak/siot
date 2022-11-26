package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type controlHandler struct {
	controlServ service.ControlService
}

func NewControlHandler(controlServ service.ControlService) *controlHandler {
	return &controlHandler{controlServ}
}

func (h *controlHandler) Create(ctx *gin.Context) {
	userId, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	body := service.CreateControlRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	body.BoxId = boxId
	body.UserId = userId
	res, err := h.controlServ.Create(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *controlHandler) FindControls(ctx *gin.Context) {
	userId, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	body := service.FindControlsRequest{BoxId: boxId, UserId: userId}
	res, err := h.controlServ.FindControls(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
