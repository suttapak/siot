package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type boxMemberRepository struct {
	db *gorm.DB
}

func NewBoxMemberRepository(db *gorm.DB) BoxMemberRepository {
	return &boxMemberRepository{db}
}

func (r *boxMemberRepository) Create(ctx context.Context, token string,
	userId, boxId uuid.UUID, canRead, canWrite bool) (boxMember *model.BoxMember, err error) {
	boxMember = &model.BoxMember{
		UserAccessToken: token,
		UserId:          userId,
		BoxId:           boxId,
	}
	err = r.db.WithContext(ctx).Create(&boxMember).Error
	if err != nil {
		return nil, err
	}
	boxMember.BoxMemberPermission = &model.BoxMemberPermission{
		BoxMemberId: boxMember.ID,
		CanRead:     canRead,
		CanWrite:    canWrite,
	}
	err = r.db.WithContext(ctx).Preload(clause.Associations).Updates(&boxMember).Error
	return boxMember, err
}

func (r *boxMemberRepository) BoxMembers(ctx context.Context, boxId uuid.UUID) (boxMember []model.BoxMember, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Where("box_id = ?", boxId).Find(&boxMember).Error
	return boxMember, err
}
