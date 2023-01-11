package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
	"github.com/suttapak/siot-backend/utils/role"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userDb struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return &userDb{db}
}

func (r *userDb) Create(ctx context.Context, email, password, firstName, lastName string) (u *model.User, err error) {
	u = &model.User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}
	err = r.db.WithContext(ctx).Create(&u).Error

	return u, err
}

func (r *userDb) FindByEmail(ctx context.Context, email string) (u *model.User, err error) {
	err = r.db.WithContext(ctx).Where("email = ? ", email).First(&u).Error
	return u, err
}

func (r *userDb) FindById(ctx context.Context, userId uuid.UUID) (u *model.User, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Preload("Setting.Notification").Where("id = ?", userId).First(&u).Error
	return u, err
}

func (r *userDb) SetRoleBasic(ctx context.Context, userId uuid.UUID) (u *model.User, err error) {
	u = &model.User{
		ID: userId,
		Roles: []*model.Role{
			{ID: role.User},
		},
	}
	err = r.db.WithContext(ctx).Where(model.User{ID: userId}).Updates(&u).Error
	return u, err
}
func (r *userDb) SetRole(ctx context.Context, userId uuid.UUID, roles ...int) (u *model.User, err error) {
	var role []*model.Role
	for _, r := range roles {
		role = append(role, &model.Role{
			ID: r,
		})
	}
	u = &model.User{
		ID:    userId,
		Roles: role,
	}
	err = r.db.Updates(&u).Error
	return u, err
}

func (r *userDb) Users(ctx context.Context) (u []model.User, err error) {
	err = r.db.WithContext(ctx).Preload(clause.Associations).Find(&u).Error
	return u, err
}

func (r *userDb) ChangePassword(ctx context.Context, uId uuid.UUID, newPwd string) (u *model.User, err error) {
	u = &model.User{Password: newPwd}
	err = r.db.WithContext(ctx).Where("id = ?", uId).Updates(&u).Error
	return u, err
}
