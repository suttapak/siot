package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type AvatarService interface {
	Update(ctx context.Context, req UpdateAvatarRequest) (res *AvatarReponse, err error)
	GenerateName(ctx context.Context, File *multipart.FileHeader) (name, url, dst string)
}
type UpdateAvatarRequest struct {
	UId    uuid.UUID `json:"userId"`
	Titile string    `json:"title"`
	Url    string    `json:"url"`
}

type AvatarReponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"userId"`
}
