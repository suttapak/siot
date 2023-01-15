package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
)

type layoutHandler struct {
	layoutServ service.LayoutService
}

func NewLayoutHandler(layoutServ service.LayoutService) *layoutHandler {
	return &layoutHandler{layoutServ}
}

func (h *layoutHandler) Update(ctx *gin.Context) {

	body := []service.UpdateLayoutRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	res, err := h.layoutServ.Update(ctx, body)
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
