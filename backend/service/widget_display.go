package service

import (
	"context"
	"time"
)

type WidgetDisplayService interface {
	Create(ctx context.Context, req *CreateWidgetDisplayRequest) (res *WidgetDisplayResponse, err error)
	Widgets(ctx context.Context) (res []WidgetDisplayResponse, err error)
	Widget(ctx context.Context, widgetId uint) (res *WidgetDisplayResponse, err error)
}

type WidgetDisplayResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DataType    string    `json:"dataType"`
}

type CreateWidgetDisplayRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	DataType    string `json:"dataType"`
}
