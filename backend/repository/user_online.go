package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
)

type UserOnlineRepository interface {
	Create(ctx context.Context, count uint) (online *model.UserOnline, err error)
	UsersOnline(ctx context.Context) (online []model.UserOnline, err error)
}
