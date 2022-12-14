package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
)

type ControlDataRepository interface {
	Crate(ctx context.Context, cId uint, data float64, label string) (c *model.ControlData, err error)
	FindByCId(ctx context.Context, cId uint) (c []model.ControlData, err error)
}
