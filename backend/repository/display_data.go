package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
)

type DisplayDataRepository interface {
	Crate(ctx context.Context, dId uint, data float64, label string) (c *model.DisplayData, err error)
}
