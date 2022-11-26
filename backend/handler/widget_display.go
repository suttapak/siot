package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type widgetDisplayHandler struct {
	widgetServ service.WidgetDisplayService
}

func NewWidgetDisplayHandler(widgetServ service.WidgetDisplayService) *widgetDisplayHandler {
	return &widgetDisplayHandler{widgetServ}
}

func (h *widgetDisplayHandler) Create(ctx *gin.Context) {
	body := service.CreateWidgetDisplayRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}

	res, err := h.widgetServ.Create(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *widgetDisplayHandler) Widgets(ctx *gin.Context) {
	res, err := h.widgetServ.Widgets(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *widgetDisplayHandler) Widget(ctx *gin.Context) {
	widgetId, err := strconv.Atoi(ctx.Param("widgetId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	res, err := h.widgetServ.Widget(ctx, uint(widgetId))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
