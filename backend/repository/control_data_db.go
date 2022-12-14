package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type controlDataDb struct {
	db *gorm.DB
}

func NewControlDataRepository(db *gorm.DB) ControlDataRepository {
	return &controlDataDb{db}
}

func (r *controlDataDb) Crate(ctx context.Context, cId uint, data float64, label string) (c *model.ControlData, err error) {
	c = &model.ControlData{
		ControlId: cId,
		Data:      data,
		Label:     label,
	}
	err = r.db.WithContext(ctx).Create(&c).Error
	return c, err
}

func (r *controlDataDb) FindByCId(ctx context.Context, cId uint) (c []model.ControlData, err error) {
	err = r.db.WithContext(ctx).Where("control_id = ?", cId).Order("id desc").Limit(4).Find(&c).Error
	return c, err
}
