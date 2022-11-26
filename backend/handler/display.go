package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"net/http"
)

type displayHandler struct {
	displayServ service.DisplayService
}

func NewDisplayHandler(displayServ service.DisplayService) *displayHandler {
	return &displayHandler{displayServ}
}

func (h *displayHandler) Create(ctx *gin.Context) {
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
	body := service.CreateDisplayRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	body.BoxId = boxId
	body.UserId = userId

	res, err := h.displayServ.Create(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *displayHandler) FindDisplays(ctx *gin.Context) {
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
	body := service.FindDisplaysRequest{BoxId: boxId, UserId: userId}

	res, err := h.displayServ.FindDisplay(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
