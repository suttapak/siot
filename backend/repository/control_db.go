package repository

import (
	"context"

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
