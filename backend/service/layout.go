package service

import (
	"context"
	"time"
)

type LayoutService interface {
	Update(ctx context.Context, req []UpdateLayoutRequest) (res *LayoutResponse, err error)
}

type UpdateLayoutRequest struct {
	ID uint   `json:"id" binding:"required"`
	I  string `json:"i"`
	X  int    `json:"x"`
	Y  int    `json:"y"`
	W  int    `json:"w"`
	H  int    `json:"h"`
}

type LayoutResponse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	I         string    `json:"i"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	W         int       `json:"w"`
	H         int       `json:"h"`
}
