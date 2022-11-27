package service

import (
	"context"
	"errors"
	"strings"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"github.com/suttapak/siot-backend/utils/role"
	"gorm.io/gorm"
)

type mqttAuthService struct {
	boxRepo    repository.BoxRepository
	canSubRepo repository.CanSubRepository
	canPubRepo repository.CanPubRepository
	userRepo   repository.UserRepository
}

func NewMqttAuthService(boxRepo repository.BoxRepository, canSubRepo repository.CanSubRepository, canPubRepo repository.CanPubRepository, userRepo repository.UserRepository) MqttAuthService {
	return &mqttAuthService{boxRepo, canSubRepo, canPubRepo, userRepo}
}

func (s *mqttAuthService) Auth(ctx context.Context, req *MqttAuthRequest) error {
	_, err := s.boxRepo.FindBoxBySecret(ctx, req.BoxId, req.Secret)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrNotFound
		}
		return errs.ErrInternalServerError
	}
	return err
}

func (s *mqttAuthService) ACLCheckI(ctx context.Context, req *MqttACLRequest) error {
	//
	_, err := s.boxRepo.FindBoxById(ctx, req.BoxId)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrNotFound
		}
		return errs.ErrInternalServerError
	}

	/*
			read       = 1
			write      = 2
		   	readWrite  = 3
		   	subscribe  = 4
	*/

	switch req.Acc {
	case 1:
		if !s.canSub(ctx, req) {
			return errs.ErrUnauthorized
		}
	case 2:
		if !s.canPub(ctx, req) {
			return errs.ErrUnauthorized
		}
	case 3:
		if !s.canPub(ctx, req) || !s.canSub(ctx, req) {
			return errs.ErrUnauthorized
		}
	case 4:
		if !s.canSub(ctx, req) {
			return errs.ErrUnauthorized
		}
	default:
		return errs.ErrBadRequest
	}

	return err
}

func (s *mqttAuthService) canSub(ctx context.Context, req *MqttACLRequest) bool {
	canSub, err := s.canSubRepo.CanSub(ctx, req.BoxId)
	if err != nil {
		return false
	}
	return s.getTopic(req) == canSub.CanSubscribe

}

func (s *mqttAuthService) canPub(ctx context.Context, req *MqttACLRequest) bool {
	canPub, err := s.canPubRepo.CanPub(ctx, req.BoxId)
	if err != nil {
		return false
	}

	return s.getTopic(req) == canPub.CanPublish

}

func (s *mqttAuthService) getTopic(req *MqttACLRequest) string {
	topic := strings.SplitN(req.Topic, "/", 2)
	return topic[0]
}

func (s *mqttAuthService) Admin(ctx context.Context, req *MqttAdminRequest) error {
	u, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		logs.Error(err)
		return errs.ErrUnauthorized
	}
	if !u.PasswordIsCorrect(req.Password) {
		logs.Error(err)
		return errs.ErrUnauthorized
	}
	for _, r := range u.Roles {
		if r.ID == role.SuperAdmin {
			return nil
		}
	}
	return errs.ErrUnauthorized
}
