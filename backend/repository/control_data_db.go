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

func (d *controlDataDb) Crate(ctx context.Context, cId uint, data float64, label string) (c *model.ControlData, err error) {
	c = &model.ControlData{
		ControlId: cId,
		Data:      data,
		Label:     label,
	}
	err = d.db.WithContext(ctx).Create(&c).Error
	return c, err
}
