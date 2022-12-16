package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type controlRepository struct {
	db *gorm.DB
}

func NewControlRepository(db *gorm.DB) ControlRepository {

	return &controlRepository{db}
}

func (r *controlRepository) Create(ctx context.Context, req *CreateControlRequest) (control *model.Control, err error) {

	control = &model.Control{
		Name:        req.Name,
		Key:         req.Key,
		Description: req.Description,
		BoxId:       req.BoxId,
		LayoutId:    req.LayoutId,
		WidgetId:    req.WidgetId,
	}
	err = r.db.WithContext(ctx).Create(&control).Error
	return control, err
}

func (r *controlRepository) FindControl(ctx context.Context, req *FindControlRequest) (control *model.Control, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("id = ? ", req.ControlId).First(&control).Error
	return control, err
}
func (r *controlRepository) FindControls(ctx context.Context, req *FindControlsRequest) (control []model.Control, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("box_id = ?", req.BoxId).Find(&control).Error
	return control, err
}

func (r *controlRepository) FindControlsByKey(ctx context.Context, boxId uuid.UUID, key string) (control []model.Control, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("key = ?", key).Find(&control).Error
	return control, err
}

func (r *controlRepository) Update(ctx context.Context, cId uint, req *UpdateControlRequest) (c *model.Control, err error) {
	c = &model.Control{
		Model: model.Model{
			ID: cId,
		},
		Name:        req.Name,
		Key:         req.Key,
		Description: req.Description,
	}
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("id = ? ", cId).Updates(&c).Error
	return c, err
}
func (r *controlRepository) Delete(ctx context.Context, cId uint) error {
	var err error
	var c model.Control
	err = r.db.WithContext(ctx).Where("id = ? ", cId).Delete(&c).Error
	return err
}
