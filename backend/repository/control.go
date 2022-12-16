package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type ControlRepository interface {
	Create(ctx context.Context, req *CreateControlRequest) (control *model.Control, err error)
	FindControl(ctx context.Context, req *FindControlRequest) (control *model.Control, err error)
	FindControls(ctx context.Context, req *FindControlsRequest) (control []model.Control, err error)
	FindControlsByKey(ctx context.Context, boxId uuid.UUID, key string) (control []model.Control, err error)
	Update(ctx context.Context, cId uint, req *UpdateControlRequest) (c *model.Control, err error)
	Delete(ctx context.Context, cId uint) error
}

type FindControlsRequest struct {
	BoxId uuid.UUID
}

type FindControlRequest struct {
	ControlId uint
}

type CreateControlRequest struct {
	Name        string
	Key         string    `json:"key"`
	Description string    `json:"description"`
	BoxId       uuid.UUID `json:"boxId"`
	LayoutId    uint
	WidgetId    uint
}

type UpdateControlRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
}
