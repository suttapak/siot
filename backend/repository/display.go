package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type DisplayRepository interface {
	Create(ctx context.Context, req *CreateDisplayRequest) (control *model.Display, err error)
	FindDisplay(ctx context.Context, req *FindDisplayRequest) (control *model.Display, err error)
	FindDisplays(ctx context.Context, req *FindDisplaysRequest) (control []model.Display, err error)
}

type FindDisplaysRequest struct {
	BoxId uuid.UUID
}

type FindDisplayRequest struct {
	DisplayId uint
}

type CreateDisplayRequest struct {
	Name        string
	Key         string    `json:"key"`
	Description string    `json:"description"`
	BoxId       uuid.UUID `json:"boxId"`
	LayoutId    uint
	WidgetId    uint
}
