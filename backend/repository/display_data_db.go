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
func (r *displayDataDb) FindByDisplayId(ctx context.Context, dId uint) (d []model.DisplayData, err error) {
	err = r.db.WithContext(ctx).Where("display_id = ?", dId).Limit(20).Order("id desc").Find(&d).Error
	return d, err
}
