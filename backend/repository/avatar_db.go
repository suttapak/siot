package repository

import (
	"github.com/suttapak/siot-backend/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type avatarRepository struct {
	db *gorm.DB
}

func NewAvatarRepository(db *gorm.DB) AvatarRepository {
	return &avatarRepository{db}
}

func (r *avatarRepository) Create(ctx context.Context, req CreateAvatarRequest) (a *model.Avatar, err error) {
	a = &model.Avatar{
		UserId: req.UId,
		Title:  req.Title,
		Url:    req.Url,
	}
	err = r.db.WithContext(ctx).Create(&a).Error
	return a, err
}
