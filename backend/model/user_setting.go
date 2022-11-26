package model

import "github.com/google/uuid"

type UserSetting struct {
	Model
	// attibute
	// fk
	UserId uuid.UUID `json:"userId"`
	// relation
	Notification UserNotification `gorm:"foreignKey:SettingId" json:"notification"`
}
