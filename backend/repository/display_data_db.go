package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type displayDataDb struct {
	db *gorm.DB
}

func NewDisplayDataRepositoryDb(db *gorm.DB) DisplayDataRepository {
	return &displayDataDb{db}
}

func (r *displayDataDb) Crate(ctx context.Context, dId uint, data float64, label string) (c *model.DisplayData, err error) {
	c = &model.DisplayData{
		DisplayId: dId,
		Data:      data,
		Label:     label,
	}
	err = r.db.WithContext(ctx).Create(&c).Error
	return c, err
}
