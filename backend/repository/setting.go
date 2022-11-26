package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type SettingRepository interface {
	Create(ctx context.Context, userId uuid.UUID) (user *model.UserSetting, err error)
}
