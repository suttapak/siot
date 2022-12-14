package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type displayDatahandler struct {
	displayDataServ service.DisplayDataService
}

func NewDisplayDataHandler(displayDataServ service.DisplayDataService) *displayDatahandler {
	return &displayDatahandler{displayDataServ}
}

func (h *displayDatahandler) Displays(ctx *gin.Context) {
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}

	dId, err := strconv.Atoi(ctx.Param("displayId"))
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	body := service.DisplaysDataRequest{
		BId: boxId,
		DId: uint(dId),
	}
	res, err := h.displayDataServ.Displays(ctx, body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
