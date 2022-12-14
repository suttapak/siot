package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DisplayDataService interface {
	Displays(ctx context.Context, req DisplaysDataRequest) (res []DisplayDataResponse, err error)
}

type DisplaysDataRequest struct {
	BId uuid.UUID `json:"boxId"`
	DId uint      `json:"displayId"`
}

type DisplayDataResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Data      float64   `json:"data"`
	Label     string    `json:"label"`
	DisplayId uint      `json:"displayId"`
}
