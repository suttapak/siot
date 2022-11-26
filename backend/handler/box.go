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

type boxHandler struct {
	boxServ service.BoxService
}

func NewBoxHandler(boxServ service.BoxService) *boxHandler {
	return &boxHandler{boxServ}
}

func (h *boxHandler) Create(ctx *gin.Context) {
	body := service.CreateBoxRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	userId, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	body.OwnerId = userId
	res, err := h.boxServ.Create(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *boxHandler) FindBoxes(ctx *gin.Context) {
	body := service.FindBoxesRequest{}
	userId, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	body.UserId = userId
	res, err := h.boxServ.FindBoxes(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *boxHandler) FindBox(ctx *gin.Context) {
	userId, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	body := service.FindBoxRequest{BoxId: boxId, UserId: userId}
	res, err := h.boxServ.FindBoxe(ctx, &body)

	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
