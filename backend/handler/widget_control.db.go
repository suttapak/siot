package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
	"net/http"
	"strconv"
)

type widgetControlHandler struct {
	widgetServ service.WidgetControlService
}

func NewWidgetControlHandler(widgetServ service.WidgetControlService) *widgetControlHandler {
	return &widgetControlHandler{widgetServ}
}

func (h *widgetControlHandler) Create(ctx *gin.Context) {
	body := service.CreateWidgetControlRequest{}
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

func (h *widgetControlHandler) Widgets(ctx *gin.Context) {
	res, err := h.widgetServ.Widgets(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *widgetControlHandler) Widget(ctx *gin.Context) {
	widgetId, err := strconv.Atoi(ctx.Param("widgetId"))
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}

	res, err := h.widgetServ.Widget(ctx, uint(widgetId))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)

}