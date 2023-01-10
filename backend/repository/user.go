package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type UserRepository interface {
	Create(ctx context.Context, email, password, firstName, lastName string) (u *model.User, err error)
	FindByEmail(ctx context.Context, email string) (u *model.User, err error)
	FindById(ctx context.Context, userId uuid.UUID) (u *model.User, err error)
	SetRoleBasic(ctx context.Context, userId uuid.UUID) (u *model.User, err error)
	SetRole(ctx context.Context, userId uuid.UUID, role ...int) (u *model.User, err error)
	Users(ctx context.Context) (u []model.User, err error)
}
