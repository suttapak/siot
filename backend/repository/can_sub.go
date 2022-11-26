package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type CanSubRepository interface {
	Create(ctx context.Context, canSub string, boxId uuid.UUID) (conSub *model.CanSubscribe, err error)
}
