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
		return tx.Limit(10)
	}).Find(&display).Error

	return display, err
}

func (r *displayRepository) FindDisplaysByKey(ctx context.Context, boxId uuid.UUID, key string) (display []model.Display, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("key = ?", key).Order("id desc").Find(&display).Error
	return display, err
}

func (r *displayRepository) FindDisplayByBoxId(ctx context.Context, bId uuid.UUID, dId uint) (c *model.Display, err error) {
	err = r.db.Where("box_id = ? AND id = ?", bId, dId).Limit(20).Order("id desc").First(&c).Error
	return c, err
}

func (r *displayRepository) Update(ctx context.Context, dId uint, req UpdateDisplayRequest) (d *model.Display, err error) {
	d = &model.Display{
		Model: model.Model{
			ID: dId,
		},
		Name:        req.Name,
		Key:         req.Key,
		Description: req.Description,
	}
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("id = ? ", dId).Updates(&d).Error
	return d, err
}
func (r *displayRepository) Delete(ctx context.Context, dId uint) error {
	var err error
	var d model.Display
	err = r.db.WithContext(ctx).Where("id = ? ", dId).Delete(&d).Error
	return err
}
