package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
)

type mqttHandler struct {
	mqttServ service.MqttAuthService
}

func NewMqttHandler(mqttServ service.MqttAuthService) *mqttHandler {
	return &mqttHandler{mqttServ}
}

func (h *mqttHandler) Auth(ctx *gin.Context) {
	body := service.MqttAuthRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	err := h.mqttServ.Auth(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, body)
}

func (h *mqttHandler) ACLCheck(ctx *gin.Context) {
	body := service.MqttACLRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	err := h.mqttServ.ACLCheckI(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, body)
}

func (h *mqttHandler) Admin(ctx *gin.Context) {
	body := service.MqttAdminRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	err := h.mqttServ.Admin(ctx, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, body)
}
