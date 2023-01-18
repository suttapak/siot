package service

import (
	"context"
	"math"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type userOnlineService struct {
	userOnlineRepo repository.UserOnlineRepository
}

func NewUserOnlineService(userOnlineRepo repository.UserOnlineRepository) UserOnlineService {
	return &userOnlineService{userOnlineRepo}
}

func (s *userOnlineService) Decrement(ctx context.Context) (res []UserOnlineResponse, err error) {
	online, err := s.userOnlineRepo.UsersOnline(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	if !math.Signbit(float64(online[0].OnlineCount)) {
		_, err = s.userOnlineRepo.Create(ctx, online[0].OnlineCount-1)
		if err != nil {
			logs.Error(err)
			return nil, errs.ErrInternalServerError
		}
	}
	online, err = s.userOnlineRepo.UsersOnline(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	res, err = utils.Recast[[]UserOnlineResponse](online)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
func (s *userOnlineService) Increment(ctx context.Context) (res []UserOnlineResponse, err error) {
	online, err := s.userOnlineRepo.UsersOnline(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	_, err = s.userOnlineRepo.Create(ctx, online[0].OnlineCount+1)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	online, err = s.userOnlineRepo.UsersOnline(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	res, err = utils.Recast[[]UserOnlineResponse](online)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
