package service

import (
	"context"

	"github.com/google/uuid"
)

type CanSubService interface {
	CanSub(ctx context.Context, boxId uuid.UUID, boxSecret string) (res *CanSubResponse, err error)
}

type CanSubResponse struct {
	CanSub string `json:"subpub"`
}
