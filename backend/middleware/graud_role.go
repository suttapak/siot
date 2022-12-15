package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
)

type graudRole struct {
	boxMemberRepo repository.BoxMemberRepository
}

func NewGraudRole(boxMemberRepo repository.BoxMemberRepository) *graudRole {
	return &graudRole{boxMemberRepo}
}

func (m *graudRole) CanWrite(ctx *gin.Context) {
	// get user id form jwt.
	// get box id form params.
	// find box member and then ckeck permission.
	// check premisstion in box.
	// if can't write reject request.
	// if can wirte continute.

	// get user id form jwt.
	userId, err := utils.UserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(err))
		return
	}
	// get box id form params.
	boxId, err := uuid.Parse(ctx.Param("boxId"))
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrBadRequest))
		return
	}
	// find box member and then ckeck permission.
	member, err := m.boxMemberRepo.BoxMember(ctx, boxId, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(handleError(errs.ErrUnauthorized))
		return
	}

	// if can't write reject request.
	if !member.BoxMemberPermission.CanWrite {
		ctx.AbortWithStatusJSON(handleError(errs.ErrUnauthorized))
		return
	}
	// if can wirte continute.
	ctx.Next()
}
