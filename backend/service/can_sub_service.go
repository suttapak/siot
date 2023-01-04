package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type canSubService struct {
	boxRepo    repository.BoxRepository
	canSubRepo repository.CanSubRepository
}

func NewCanSubService(boxRepo repository.BoxRepository,
	canSubRepo repository.CanSubRepository) CanSubService {
	return &canSubService{boxRepo, canSubRepo}
}

func (s *canSubService) CanSub(ctx context.Context, boxId uuid.UUID, boxSecret string) (res *CanSubResponse, err error) {

	_, err = s.boxRepo.FindBoxBySecret(ctx, boxId, boxSecret)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrUnauthorized
	}

	canSub, err := s.canSubRepo.CanSub(ctx, boxId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrUnauthorized
	}
	res = &CanSubResponse{
		CanSub: canSub.CanSubscribe,
	}
	return res, err
}
