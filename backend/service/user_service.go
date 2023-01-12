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

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) FindUser(ctx context.Context, req *FindUserRequest) (res *UserResponse, err error) {
	user, err := u.userRepo.FindById(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrNotFound
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*UserResponse](user)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, nil
}

func (u *userService) FindUsers(ctx context.Context) (res []UserResponse, err error) {
	user, err := u.userRepo.Users(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]UserResponse](user)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

func (u *userService) AddRoles(ctx context.Context, req *AddRolesRequest) (res *UserResponse, err error) {
	// check user exist ?
	user, err := u.userRepo.FindById(ctx, req.UserId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrBadRequest
	}
	roleExist := false
	// check role exist ?
	for _, v := range user.Roles {
		if v.ID == req.Role {
			roleExist = true
		}
	}
	if roleExist {
		logs.Error("roles is exist")
		return nil, errs.ErrBadRequest
	}
	res, err = utils.Recast[*UserResponse](user)
	if err != nil {
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
