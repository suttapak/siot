package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type CanSubRepository interface {
	Create(ctx context.Context, canSub string, boxId uuid.UUID) (conSub *model.CanSubscribe, err error)
	CanSub(ctx context.Context, boxId uuid.UUID) (canSub *model.CanSubscribe, err error)
	GetCanSubByTopic(ctx context.Context, topic string) (canSub *model.CanSubscribe, err error)
}
