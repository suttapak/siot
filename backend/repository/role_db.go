package repository

import (
	"github.com/suttapak/siot-backend/model"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	roles := []model.Role{
		{
			ID:              RoleUser,
			Name:            "user",
			PermissionState: 1,
			DisplayName:     "User",
			Description:     "User",
		},
		{
			ID:              RoleAdmin,
			Name:            "admin",
			PermissionState: 3,
			DisplayName:     "Admin",
			Description:     "Admin",
		},
		{
			ID:              RoleSuperAdmin,
			Name:            "superAdmin",
			PermissionState: 1,
			DisplayName:     "Super Admin",
			Description:     "Super Admin",
		},
	}

	err := db.Create(&roles).Error
	if err != nil {
		logs.Error(err)
	}

	return &roleRepository{db: db}
}
