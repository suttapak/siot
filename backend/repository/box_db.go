package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type boxRepository struct {
	db *gorm.DB
}

func NewBoxRepository(db *gorm.DB) BoxRepository {
	return &boxRepository{db}
}

func (b *boxRepository) Create(ctx context.Context, name, description string, ownerId uuid.UUID) (box *model.Box, err error) {
	box = &model.Box{
		Name:        name,
		Description: description,
		OwnerId:     ownerId,
	}
	err = b.db.WithContext(ctx).Create(&box).Error
	return box, err
}
func (b *boxRepository) FindAll(ctx context.Context, userId uuid.UUID) (boxes []model.Box, err error) {
	err = b.db.WithContext(ctx).Preload(clause.Associations).
		Preload("Members.BoxMemberPermission").Where("owner_id = ?", userId).Order("created_at ASC").Find(&boxes).Error
	return boxes, err
}

func (b *boxRepository) FindBox(ctx context.Context, boxId, userId uuid.UUID) (box *model.Box, err error) {
	err = b.db.WithContext(ctx).Preload(clause.Associations).Preload("Members.BoxMemberPermission").
		Where("id = ? AND owner_id = ?", boxId, userId).First(&box).Error
	return box, err
}

func (b *boxRepository) FindBoxBySecret(ctx context.Context, boxId uuid.UUID, secret string) (box *model.Box, err error) {
	err = b.db.WithContext(ctx).Preload(clause.Associations).Where(model.Box{ID: boxId, BoxSecret: model.BoxSecret{Secret: secret}}).First(&box).Error
	return box, err
}

func (b *boxRepository) FindBoxById(ctx context.Context, boxId uuid.UUID) (box *model.Box, err error) {
	err = b.db.WithContext(ctx).Where("id = ?", boxId).Preload(clause.Associations).First(&box).Error
	return box, err
}

func (b *boxRepository) UpdateBox(ctx context.Context, req UpdateBoxRequest, bId uuid.UUID) (box *model.Box, err error) {
	box = &model.Box{
		ID:          bId,
		Name:        req.Name,
		Description: req.Description,
	}
	err = b.db.WithContext(ctx).Updates(&box).Error
	return box, err
}

func (b *boxRepository) DeleteBox(ctx context.Context, bId uuid.UUID) error {
	err := b.db.WithContext(ctx).Delete(&model.Box{}, bId).Error
	return err
}

func (b *boxRepository) FindBoxByMember(ctx context.Context, userId uuid.UUID) (box []model.Box, err error) {
	box = []model.Box{}
	bm := []model.BoxMember{}
	if err := b.db.WithContext(ctx).Where("user_id = ?", userId).Find(&bm).Error; err != nil {
		return nil, err
	}
	var rx []uuid.UUID
	for _, v := range bm {
		rx = append(rx, v.BoxId)
	}
	if err := b.db.WithContext(ctx).Preload(clause.Associations).Where("id IN ?", rx).Find(&box).Error; err != nil {
		return nil, err
	}
	return box, err
}
