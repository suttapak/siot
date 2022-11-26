package service

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type BoxMemberService interface {
	BoxMembers(ctx context.Context, boxId uuid.UUID) (res []BoxMemberResponse, err error)
	AddMember(ctx context.Context, boxId uuid.UUID, req *AddBoxMemberRequest) (res *BoxMemberResponse, err error)
}

type AddBoxMemberRequest struct {
	UserEmail string `json:"userEmail" binding:"required"`
	CanRead   bool   `json:"canRead" binding:"required"`
	CanWrite  bool   `json:"canWrite" binding:"required"`
}

type BoxMemberResponse struct {
	ID                  uint      `json:"id"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	UserAccessToken     string    `json:"userAccessToken"`
	UserId              uuid.UUID `json:"userId"`
	BoxId               uuid.UUID `json:"boxId"`
	BoxMemberPermission struct {
		ID          uint      `json:"id"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		CanRead     bool      `json:"canRead"`
		CanWrite    bool      `json:"canWrite"`
		BoxMemberId uint      `json:"boxMemberId"`
	} `json:"boxMemberPermission"`
	User User `json:"user"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}
