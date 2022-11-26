package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type BoxRepository interface {
	Create(ctx context.Context, name, description string, ownerId uuid.UUID) (box *model.Box, err error)
	FindAll(ctx context.Context, userId uuid.UUID) (boxes []model.Box, err error)
	FindBox(ctx context.Context, boxId, userId uuid.UUID) (box *model.Box, err error)
}
