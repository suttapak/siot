package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"net/http"
)

type boxMemberHandler struct {
	boxMemberServ service.BoxMemberService
}

func NewBoxMemberHandler(boxMemberServ service.BoxMemberService) *boxMemberHandler {
	return &boxMemberHandler{boxMemberServ}
}

func (b *boxMemberHandler) BoxMembers(ctx *gin.Context) {
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	res, err := b.boxMemberServ.BoxMembers(ctx, boxId)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (b *boxMemberHandler) AddMember(ctx *gin.Context) {
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}

	body := service.AddBoxMemberRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	res, err := b.boxMemberServ.AddMember(ctx, boxId, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
