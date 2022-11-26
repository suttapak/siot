package service

import (
	"context"
	"time"
)

type WidgetControlService interface {
	Create(ctx context.Context, req *CreateWidgetControlRequest) (res *WidgetControlResponse, err error)
	Widgets(ctx context.Context) (res []WidgetControlResponse, err error)
	Widget(ctx context.Context, widgetId uint) (res *WidgetControlResponse, err error)
}

type WidgetControlResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DataType    string    `json:"dataType"`
}

type CreateWidgetControlRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" `
	DataType    string `json:"dataType"`
}
