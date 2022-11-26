package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type boxSecretRepository struct {
	db *gorm.DB
}

func NewBoxSecretRepository(db *gorm.DB) BoxSecretRepository {
	return &boxSecretRepository{db}
}

func (b *boxSecretRepository) Create(ctx context.Context, secret string, boxId uuid.UUID) (boxSecret *model.BoxSecret, err error) {
	boxSecret = &model.BoxSecret{
		Secret: secret,
		BoxId:  boxId,
	}
	err = b.db.WithContext(ctx).Create(&boxSecret).Error
	return boxSecret, err
}
