package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type settingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

func (r *settingRepository) Create(ctx context.Context, userId uuid.UUID) (user *model.UserSetting, err error) {

	user = &model.UserSetting{
		UserId: userId,
	}

	err = r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	notification := &model.UserNotification{
		NotificationState: false,
		SettingId:         user.ID,
	}
	err = r.db.Create(&notification).Error
	return user, err
}
