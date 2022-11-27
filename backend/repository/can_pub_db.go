package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type canPubRepository struct {
	db *gorm.DB
}

func NewCanPubRepository(db *gorm.DB) CanPubRepository {
	return &canPubRepository{db}
}

func (c *canPubRepository) Create(ctx context.Context, canPubs string, boxId uuid.UUID) (canPub *model.CanPublish, err error) {
	canPub = &model.CanPublish{CanPublish: canPubs, BoxId: boxId}
	err = c.db.WithContext(ctx).Create(&canPub).Error
	return canPub, err
}

func (c *canPubRepository) CanPub(ctx context.Context, boxId uuid.UUID) (canPub *model.CanPublish, err error) {
	err = c.db.WithContext(ctx).Where("box_id = ? ", boxId).First(&canPub).Error
	return canPub, err
}
