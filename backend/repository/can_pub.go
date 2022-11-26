package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type CanPubRepository interface {
	Create(ctx context.Context, canPubs string, boxId uuid.UUID) (canPub *model.CanPublish, err error)
	CanPub(ctx context.Context, boxId uuid.UUID) (canPub *model.CanPublish, err error)
}
