package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type boxService struct {
	conf          *config.Configs
	boxRepo       repository.BoxRepository
	boxMemberRepo repository.BoxMemberRepository
	boxSecretRepo repository.BoxSecretRepository
	canSub        repository.CanSubRepository
	canPub        repository.CanPubRepository
}

func NewBoxService(conf *config.Configs, boxRepo repository.BoxRepository, boxMemberRepo repository.BoxMemberRepository,
	boxSecretRepo repository.BoxSecretRepository, canSub repository.CanSubRepository, canPub repository.CanPubRepository) BoxService {
	return &boxService{conf, boxRepo, boxMemberRepo, boxSecretRepo, canSub, canPub}
}

func (s *boxService) Create(ctx context.Context, req *CreateBoxRequest) (res *BoxResponse, err error) {
	// create box
	box, err := s.boxRepo.Create(ctx, req.Name, req.Description, req.OwnerId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	_ = box
	// add member
	claims := jwt.MapClaims{
		"BoxId": box.ID,
	}
	// Create token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	token, err := t.SignedString([]byte(s.conf.JWT.Secret))
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	_, err = s.boxMemberRepo.Create(ctx, token, req.OwnerId, box.ID, true, true)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// create box secret
	secret := randStringRunes(20)
	_, err = s.boxSecretRepo.Create(ctx, secret, box.ID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// create can-pub and can-sub
	canSubAndPub := fmt.Sprintf("%v-%v", strings.ReplaceAll(box.Name, " ", "-"), strings.Split(box.ID.String(), "-")[0])
	_, err = s.canSub.Create(ctx, canSubAndPub, box.ID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	_, err = s.canPub.Create(ctx, canSubAndPub, box.ID)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	box, err = s.boxRepo.FindBox(ctx, box.ID, req.OwnerId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*BoxResponse](box)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
func (s *boxService) FindBoxes(ctx context.Context, req *FindBoxesRequest) (res []BoxResponse, err error) {
	boxes, err := s.boxRepo.FindAll(ctx, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrNotFound
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]BoxResponse](boxes)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

func (s *boxService) FindBoxe(ctx context.Context, req *FindBoxRequest) (res *BoxResponse, err error) {
	box, err := s.boxRepo.FindIsMember(ctx, req.BoxId, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrNotFound
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*BoxResponse](box)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

func (s *boxService) Update(ctx context.Context, uId, bId uuid.UUID, req UpdateBoxRequest) (res *BoxResponse, err error) {
	_, err = s.boxRepo.FindIsMember(ctx, bId, uId)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUnauthorized
		}
		return nil, errs.ErrInternalServerError
	}
	box, err := s.boxRepo.UpdateBox(ctx, repository.UpdateBoxRequest(req), bId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*BoxResponse](box)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
func (s *boxService) Delete(ctx context.Context, uId, bId uuid.UUID) error {
	var err error
	_, err = s.boxRepo.FindIsMember(ctx, bId, uId)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.ErrUnauthorized
		}
		return errs.ErrInternalServerError
	}
	err = s.boxRepo.DeleteBox(ctx, bId)
	if err != nil {
		return errs.ErrInternalServerError
	}
	return err
}

func (s *boxService) FindBoxByMember(ctx context.Context, uId uuid.UUID) (res []BoxResponse, err error) {
	box, err := s.boxRepo.FindBoxByMember(ctx, uId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]BoxResponse](box)
	if err != nil {
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
