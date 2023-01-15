package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
)

type LayoutRepository interface {
	Create(ctx context.Context, req *CreateLayoutRequest) (layout *model.Layout, err error)
	Update(ctx context.Context, req []UpdateLayoutRequest) (layout *model.Layout, err error)
}

type UpdateLayoutRequest struct {
	ID uint   `json:"id"`
	I  string `json:"i"`
	X  int    `json:"x"`
	Y  int    `json:"y"`
	W  int    `json:"w"`
	H  int    `json:"h"`
}

type CreateLayoutRequest struct {
	I string `json:"i"`
	X int    `json:"x"`
	Y int    `json:"y"`
	W int    `json:"w"`
	H int    `json:"h"`
}
