package repository

import (
	"context"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type layoutRepository struct {
	db *gorm.DB
}

func NewLayoutRepository(db *gorm.DB) LayoutRepository {
	return &layoutRepository{db}
}

func (r *layoutRepository) Create(ctx context.Context, req *CreateLayoutRequest) (layout *model.Layout, err error) {
	layout = &model.Layout{I: req.I, X: req.X, Y: req.Y, H: req.H, W: req.W}
	err = r.db.WithContext(ctx).Create(&layout).Error
	return layout, err
}
