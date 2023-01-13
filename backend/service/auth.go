package service

import (
	"context"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, req *LoginRequest) (res *LoginRespose, err error)
	Register(ctx context.Context, req *RegisterRequest) (res *RegisterResponse, err error)
	ChangePassword(ctx context.Context, uId uuid.UUID, req *ChangePasswordRequest) error
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// login response
type LoginRespose struct {
	AccessToken string `json:"accessToken"`
}

// register request
type RegisterRequest struct {
	Email     string `json:"email"  binding:"required,email"`
	Password  string `json:"password"  binding:"required,min=6,max=32"`
	FirstName string `json:"firstName"  binding:"required"`
	LastName  string `json:"lastName"  binding:"required"`
}

// register response
type RegisterResponse struct {
	AccessToken string `json:"accessToken"`
}
