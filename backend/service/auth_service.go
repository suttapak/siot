package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type authService struct {
	avatarRepo  repository.AvatarRepository
	userRepo    repository.UserRepository
	conf        *config.Configs
	settingRepo repository.SettingRepository
}

func NewAuthService(avatarRepo repository.AvatarRepository, userRepo repository.UserRepository, conf *config.Configs, settingRepository repository.SettingRepository) AuthService {
	return &authService{avatarRepo, userRepo, conf, settingRepository}
}

func (s *authService) Login(ctx context.Context, req *LoginRequest) (res *LoginRespose, err error) {
	u, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternalServerError
	}
	if !u.PasswordIsCorrect(req.Password) {
		logs.Error("err password incorrect")
		return nil, errs.ErrUnauthorized
	}
	// TODO : change to custom claims

	claims := jwt.MapClaims{
		"email":     u.Email,
		"userId":    u.ID,
		"ExpiresAt": time.Now().Add(time.Hour * time.Duration(s.conf.JWT.TTLHour)).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(s.conf.JWT.Secret))
	if err != nil {
		return nil, errs.ErrBadGateway
	}
	res = &LoginRespose{
		AccessToken: t,
	}
	return res, err
}
func (s *authService) Register(ctx context.Context, req *RegisterRequest) (res *RegisterResponse, err error) {
	// check out in database
	// if user is not exist
	// first user is create set role are admin user and super admin

	users, _ := s.userRepo.Users(ctx)
	u, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrInternalServerError
		}
	}
	if u.EmailIsExist() {
		logs.Error(err)
		return nil, errs.ErrUnauthorized
	}

	u, err = s.userRepo.Create(ctx, req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	_, err = s.userRepo.SetRoleBasic(ctx, u.ID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// if first user set role user admin and super admin
	if len(users) <= 0 {
		s.userRepo.SetRole(ctx, u.ID, 1, 2, 3)
	}
	_, err = s.avatarRepo.Create(ctx, repository.CreateAvatarRequest{UId: u.ID})
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// TODO : create setting user
	_, err = s.settingRepo.Create(ctx, u.ID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// TODO : change to custom claims

	claims := jwt.MapClaims{
		"email":     u.Email,
		"userId":    u.ID,
		"ExpiresAt": time.Now().Add(time.Hour * time.Duration(s.conf.JWT.TTLHour)).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(s.conf.JWT.Secret))
	if err != nil {
		return nil, errs.ErrBadGateway
	}
	res = &RegisterResponse{
		AccessToken: t,
	}
	return res, err
}
