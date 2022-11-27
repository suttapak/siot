package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type canSubRepository struct {
	db *gorm.DB
}

func NewCanSubRepository(db *gorm.DB) CanSubRepository {
	return &canSubRepository{db}
}

func (r *canSubRepository) Create(ctx context.Context, canSubs string, boxId uuid.UUID) (canSub *model.CanSubscribe, err error) {
	canSub = &model.CanSubscribe{
		CanSubscribe: canSubs,
		BoxId:        boxId,
	}
	err = r.db.WithContext(ctx).Create(&canSub).Error
	return canSub, err
}

func (r *canSubRepository) CanSub(ctx context.Context, boxId uuid.UUID) (canSub *model.CanSubscribe, err error) {
	err = r.db.WithContext(ctx).Where("box_id = ?", boxId).First(&canSub).Error
	return canSub, err
}
