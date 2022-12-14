package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type AvatarRepository interface {
	Create(ctx context.Context, req CreateAvatarRequest) (a *model.Avatar, err error)
}

type CreateAvatarRequest struct {
	UId   uuid.UUID `json:"userId"`
	Title string
	Url   string
}
