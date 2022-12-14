package service

import (
	"context"
	"errors"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type displayDataService struct {
	displayRepo repository.DisplayRepository
	dataRepo    repository.DisplayDataRepository
}

func NewDisplayDataService(displayRepo repository.DisplayRepository, dataRepo repository.DisplayDataRepository) DisplayDataService {
	return &displayDataService{displayRepo, dataRepo}
}

func (s *displayDataService) Displays(ctx context.Context, req DisplaysDataRequest) (res []DisplayDataResponse, err error) {
	_, err = s.displayRepo.FindDisplayByBoxId(ctx, req.BId, req.DId)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternalServerError
	}
	d, err := s.dataRepo.FindByDisplayId(ctx, req.DId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]DisplayDataResponse](d)
	return res, err
}
