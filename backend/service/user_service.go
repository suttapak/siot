package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) FindUser(ctx context.Context, req *FindUserRequest) (res *FindUserResponse, err error) {
	user, err := u.userRepo.FindById(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrNotFound
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*FindUserResponse](user)
	fmt.Println(res.Avatar)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, nil
}
