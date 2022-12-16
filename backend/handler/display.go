package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
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

func (h *displayHandler) Update(ctx *gin.Context) {
	cId, err := strconv.Atoi(ctx.Param("displayId"))
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	body := service.UpdateDisplayRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	res, err := h.displayServ.Update(ctx, uint(cId), &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *displayHandler) Delete(ctx *gin.Context) {
	cId, err := strconv.Atoi(ctx.Param("displayId"))
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	err = h.displayServ.Delete(ctx, uint(cId))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	ctx.JSON(http.StatusCreated, ResponseOk)
}
