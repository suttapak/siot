package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type BoxRepository interface {
	Create(ctx context.Context, name, description string, ownerId uuid.UUID) (box *model.Box, err error)
	FindAll(ctx context.Context, userId uuid.UUID) (boxes []model.Box, err error)
	FindBoxBySecret(ctx context.Context, boxId uuid.UUID, secret string) (box *model.Box, err error)
	FindBox(ctx context.Context, boxId, userId uuid.UUID) (box *model.Box, err error)
	FindBoxById(ctx context.Context, boxId uuid.UUID) (box *model.Box, err error)
	UpdateBox(ctx context.Context, req UpdateBoxRequest, bId uuid.UUID) (box *model.Box, err error)
	DeleteBox(ctx context.Context, bId uuid.UUID) error
}

type UpdateBoxRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
