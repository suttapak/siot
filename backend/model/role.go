package model

import "time"

type Role struct {
	// common
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	// attribute
	Name            string `json:"name"`
	PermissionState int    `json:"permissionState"`
	DisplayName     string `json:"displayName"`
	Description     string `json:"description"`
	// fk
	Users []*User `gorm:"many2many:user_roles;" json:"users"`
}
