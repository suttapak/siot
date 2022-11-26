package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ControlService interface {
	Create(ctx context.Context, req *CreateControlRequest) (res *ControlResponse, err error)
	FindControls(ctx context.Context, req *FindControlsRequest) (res []ControlResponse, err error)
}

type FindControlsRequest struct {
	UserId uuid.UUID `json:"userId"`
	BoxId  uuid.UUID `json:"boxId"`
}

type CreateControlRequest struct {
	UserId      uuid.UUID
	Name        string    `json:"name" binding:"required"`
	Key         string    `json:"key" binding:"required,min=2,max=6"`
	Description string    `json:"description"`
	BoxId       uuid.UUID `json:"BoxId"`
	Layout      struct {
		I string `json:"i" binding:"required"`
		X int    `json:"x" binding:"number"`
		Y int    `json:"y" binding:"number"`
		W int    `json:"w" binding:"required"`
		H int    `json:"h" binding:"required"`
	} `json:"layout" binding:"required"`
	Widget struct {
		ID uint `json:"id" binding:"required"`
	} `json:"widget" binding:"required"`
}

type ControlResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name" `
	Key         string    `json:"key"`
	Description string    `json:"description"`
	BoxId       uuid.UUID `json:"BoxId"`
	LayoutId    uint      `json:"layoutId"`
	WidgetId    uint      `json:"widgetId"`
	ControlData []struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		XNumber   int       `json:"XNumber"`
		YNumber   int       `json:"YNumber"`
		XString   string    `json:"XString"`
		YString   string    `json:"YString"`
		ControlId uint      `json:"controlId"`
	} ` json:"controlData"`
	Widget struct {
		ID          uint      `json:"id"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		DataType    string    `json:"dataType"`
	} `json:"widget"`
	Layout struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		I         string    `json:"i"`
		X         int       `json:"x"`
		Y         int       `json:"y"`
		W         int       `json:"w"`
		H         int       `json:"h"`
	} `json:"layout"`
}
