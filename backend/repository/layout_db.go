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

func (r *layoutRepository) Update(ctx context.Context, req []UpdateLayoutRequest) (layout *model.Layout, err error) {
	body := []model.Layout{}
	for _, l := range req {
		body = append(body, model.Layout{
			Model: model.Model{ID: l.ID},
			I:     l.I,
			X:     l.X,
			Y:     l.Y,
			W:     l.W,
			H:     l.H,
		})
	}
	err = r.db.WithContext(ctx).Model(&model.Layout{}).Save(&body).Error
	return layout, err
}
