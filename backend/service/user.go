package service

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	FindUser(ctx context.Context, req *FindUserRequest) (res *UserResponse, err error)
	FindUsers(ctx context.Context) (res []UserResponse, err error)
	AddRoles(ctx context.Context, req *AddRolesRequest) (res *UserResponse, err error)
}

type AddRolesRequest struct {
	UserId uuid.UUID `json:"userId"`
	Role   int       `json:"role"`
}

type FindUserRequest struct {
	UserId uuid.UUID
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	SettingId uint      `json:"settingId"`
	Avatar    struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Title     string    `json:"title"`
		Url       string    `json:"url"`
		UserId    uuid.UUID `json:"userId"`
	} `json:"avatar"`
	Roles []struct {
		ID              int       `json:"id"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
		Name            string    `json:"name"`
		PermissionState int       `json:"permissionState"`
		DisplayName     string    `json:"displayName"`
		Description     string    `json:"description"`
	} `json:"roles"`
	Box []struct {
		ID          uuid.UUID `json:"id"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		OwnerId     uuid.UUID `json:"ownerId"`
	} `json:"box"`
	Setting struct {
		ID           uint      `json:"id"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
		UserId       uuid.UUID `json:"userId"`
		Notification struct {
			ID        uint      `json:"id"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
			// attbute
			NotificationState bool `json:"notificationState"`
			// fk
			SettingId uint `json:"settingId"`
		} `json:"notification"`
	} `json:"setting"`
}
