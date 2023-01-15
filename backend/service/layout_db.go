package service

import (
	"context"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type layoutService struct {
	layoutRepo repository.LayoutRepository
}

func NewLayoutService(layoutRepo repository.LayoutRepository) LayoutService {
	return &layoutService{layoutRepo}
}

func (s *layoutService) Update(ctx context.Context, req []UpdateLayoutRequest) (res *LayoutResponse, err error) {
	body := []repository.UpdateLayoutRequest{}
	for _, r := range req {
		body = append(body, repository.UpdateLayoutRequest{
			ID: r.ID,
			I:  r.I,
			X:  r.X,
			W:  r.W,
			H:  r.H,
			Y:  r.Y,
		})
	}
	layout, err := s.layoutRepo.Update(ctx, body)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*LayoutResponse](layout)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
