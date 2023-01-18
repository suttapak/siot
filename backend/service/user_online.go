package service

import (
	"context"
	"time"
)

type UserOnlineService interface {
	Decrement(ctx context.Context) (res []UserOnlineResponse, err error)
	Increment(ctx context.Context) (res []UserOnlineResponse, err error)
}

type UserOnlineResponse struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	OnlineCount uint      `json:"onLineCount"`
}
