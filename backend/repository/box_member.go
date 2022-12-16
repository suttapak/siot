package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type BoxMemberRepository interface {
	Create(ctx context.Context, token string,
		userId, boxId uuid.UUID, canRead, canWrite bool) (boxMember *model.BoxMember, err error)
	BoxMembers(ctx context.Context, boxId uuid.UUID) (boxMember []model.BoxMember, err error)
	BoxMember(ctx context.Context, bId, uId uuid.UUID) (b *model.BoxMember, err error)
}
