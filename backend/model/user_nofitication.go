package model

type UserNotification struct {
	Model
	// attbute
	NotificationState bool `gorm:"default:false" json:"notificationState"`
	// fk
	SettingId uint `json:"settingId"`
	// relation
}
