package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
	"strings"
)

type mqttAuthService struct {
	boxRepo    repository.BoxRepository
	canSubRepo repository.CanSubRepository
	canPubRepo repository.CanPubRepository
}

func NewMqttAuthService(boxRepo repository.BoxRepository, canSubRepo repository.CanSubRepository, canPubRepo repository.CanPubRepository) MqttAuthService {
	return &mqttAuthService{boxRepo, canSubRepo, canPubRepo}
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

	fmt.Println(req)

	switch req.Acc {
	case 1:
		if !s.canSub(ctx, req) {
			return errs.ErrUnauthorized
		}
		break
	case 2:
		if !s.canPub(ctx, req) {
			return errs.ErrUnauthorized
		}
		break
	case 3:
		if !s.canPub(ctx, req) || !s.canPub(ctx, req) {
			return errs.ErrUnauthorized
		}
		break
	case 4:
		if !s.canSub(ctx, req) {
			return errs.ErrUnauthorized
		}
		break
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
	fmt.Println(s.getTopic(req), canSub.CanSubscribe)
	return s.getTopic(req) == canSub.CanSubscribe

}

func (s *mqttAuthService) canPub(ctx context.Context, req *MqttACLRequest) bool {
	canPub, err := s.canPubRepo.CanPub(ctx, req.BoxId)
	if err != nil {
		return false
	}
	fmt.Println(s.getTopic(req), canPub.CanPublish)

	return s.getTopic(req) == canPub.CanPublish

}

func (s *mqttAuthService) getTopic(req *MqttACLRequest) string {
	topic := strings.SplitN(req.Topic, "/", 2)
	return topic[0]
}
