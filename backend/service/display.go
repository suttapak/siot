package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DisplayService interface {
	Create(ctx context.Context, req *CreateDisplayRequest) (res *DisplayResponse, err error)
	FindDisplay(ctx context.Context, req *FindDisplaysRequest) (res []DisplayResponse, err error)
	Update(ctx context.Context, dId uint, req *UpdateDisplayRequest) (res *DisplayResponse, err error)
	Delete(ctx context.Context, dId uint) error
}

type FindDisplaysRequest struct {
	BoxId  uuid.UUID
	UserId uuid.UUID
}

type CreateDisplayRequest struct {
	UserId      uuid.UUID
	Name        string    `json:"name" binding:"required"`
	Key         string    `json:"key" binding:"required,min=2,max=6"`
	Description string    `json:"description"`
	BoxId       uuid.UUID `json:"BoxId"`
	Layout      struct {
		I string `json:"i" binding:"required"`
		X int    `json:"x"`
		Y int    `json:"y"`
		W int    `json:"w" binding:"required"`
		H int    `json:"h" binding:"required"`
	} `json:"layout" binding:"required"`
	Widget struct {
		ID uint `json:"id" binding:"required"`
	} `json:"widget" binding:"required"`
}

type UpdateDisplayRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
}

type DisplayResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name" `
	Key         string    `json:"key"`
	Description string    `json:"description"`
	BoxId       uuid.UUID `json:"BoxId"`
	LayoutId    uint      `json:"layoutId"`
	WidgetId    uint      `json:"widgetId"`
	DisplayData []struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Data      float64   `json:"data"`
		Label     string    `json:"label"`
		DisplayId uint      `json:"displayId"`
	} ` json:"displayData"`
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
