package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type displayRepository struct {
	db *gorm.DB
}

func NewDisplayRepository(db *gorm.DB) DisplayRepository {
	return &displayRepository{db}
}

func (r *displayRepository) Create(ctx context.Context, req *CreateDisplayRequest) (display *model.Display, err error) {

	display = &model.Display{
		Name:        req.Name,
		Key:         req.Key,
		Description: req.Description,
		BoxId:       req.BoxId,
		LayoutId:    req.LayoutId,
		WidgetId:    req.WidgetId,
	}

	err = r.db.WithContext(ctx).Create(&display).Error

	return display, err

}
func (r *displayRepository) FindDisplay(ctx context.Context, req *FindDisplayRequest) (display *model.Display, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where(model.Display{Model: model.Model{ID: req.DisplayId}}).First(&display).Error

	return display, err
}
func (r *displayRepository) FindDisplays(ctx context.Context, req *FindDisplaysRequest) (display []model.Display, err error) {

	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("box_id = ? ", req.BoxId).Preload("DisplayData", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("id desc")
	}).Find(&display).Error

	return display, err
}

func (r *displayRepository) FindDisplaysByKey(ctx context.Context, boxId uuid.UUID, key string) (display []model.Display, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("key = ?", key).Find(&display).Error
	return display, err
}
