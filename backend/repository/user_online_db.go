package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type userOnlineRepository struct {
	db *gorm.DB
}

func NewUserOnlineRepository(db *gorm.DB) UserOnlineRepository {
	temp := model.UserOnline{
		Model: model.Model{
			ID: 1,
		},
		OnlineCount: 0,
	}
	db.Create(&temp)
	return &userOnlineRepository{db}
}

func (r *userOnlineRepository) Create(ctx context.Context, count uint) (online *model.UserOnline, err error) {
	online = &model.UserOnline{
		OnlineCount: count,
	}
	err = r.db.WithContext(ctx).Create(&online).Error
	return online, err
}
func (r *userOnlineRepository) UsersOnline(ctx context.Context) (online []model.UserOnline, err error) {
	err = r.db.WithContext(ctx).Order("id DESC").Limit(40).Find(&online).Error
	return online, err
}
