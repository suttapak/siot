package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type boxMemberService struct {
	userRepo      repository.UserRepository
	boxMemberRepo repository.BoxMemberRepository
}

func NewBoxMemberService(userRepo repository.UserRepository, boxMemberRepo repository.BoxMemberRepository) BoxMemberService {
	return &boxMemberService{userRepo, boxMemberRepo}
}

func (s *boxMemberService) BoxMembers(ctx context.Context, boxId uuid.UUID) (res []BoxMemberResponse, err error) {
	boxMember, err := s.boxMemberRepo.BoxMembers(ctx, boxId)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]BoxMemberResponse](boxMember)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
func (s *boxMemberService) AddMember(ctx context.Context, boxId uuid.UUID, req *AddBoxMemberRequest) (res *BoxMemberResponse, err error) {
	user, err := s.userRepo.FindByEmail(ctx, req.UserEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, errs.ErrInternalServerError
	}

	boxMember, err := s.boxMemberRepo.Create(ctx, "", user.ID, boxId, req.CanRead, req.CanWrite)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*BoxMemberResponse](boxMember)
	return res, err
}
