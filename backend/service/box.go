package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type BoxService interface {
	Create(ctx context.Context, req *CreateBoxRequest) (res *BoxResponse, err error)
	FindBoxes(ctx context.Context, req *FindBoxesRequest) (res []BoxResponse, err error)
	FindBoxe(ctx context.Context, req *FindBoxRequest) (res *BoxResponse, err error)
	Update(ctx context.Context, uId, bId uuid.UUID, req UpdateBoxRequest) (res *BoxResponse, err error)
	Delete(ctx context.Context, uId, bId uuid.UUID) error
}

type CreateBoxRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" `
	OwnerId     uuid.UUID `json:"ownerId"`
}

type FindBoxesRequest struct {
	UserId uuid.UUID
}

type FindBoxRequest struct {
	BoxId  uuid.UUID
	UserId uuid.UUID
}

type UpdateBoxRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BoxResponse struct {
	ID          uuid.UUID ` json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     uuid.UUID `json:"ownerId"`
	Members     []struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		// attribute
		UserAccessToken string `json:"userAccessToken"`
		// fk
		UserId uuid.UUID `json:"userId"`
		BoxId  uuid.UUID `json:"BoxId"`
		// relation
		BoxMemberPermission *struct {
			ID        uint      `json:"id"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
			// attribute
			CanRead     bool `json:"canRead"`
			CanWrite    bool `json:"canWrite"`
			BoxMemberId uint `json:"boxMemberId"`
		} `gorm:"foreignKey:BoxMemberId" json:"boxMemberPermission"`
	} `gorm:"foreignKey:BoxId" json:"members"`
	BoxSecret struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Secret    string    `json:"secret"`
		// fk
		BoxId uuid.UUID `json:"BoxId"`
	} `gorm:"foreignKey:BoxId" json:"boxSecret"`
	CanSub struct {
		ID           uint      `json:"id"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
		CanSubscribe string    `json:"canSubscribe"`
		// fk
		BoxId uuid.UUID `json:"BoxId"`
	} `gorm:"foreignKey:BoxId" json:"canSub"`
	CanPub struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		// attribute
		CanPublish string `json:"canPublish"`
		// fk
		BoxId uuid.UUID `json:"BoxId"`
	} `gorm:"foreignKey:BoxId" json:"canPub"`
}
