package model

import "github.com/google/uuid"

type Avatar struct {
	Model
	Title  string    `json:"title" gorm:"default:'default avatar image at siot'"`
	Url    string    `json:"url" gorm:"default:'/asset/images/siot-avatar.png'"`
	UserId uuid.UUID `json:"userId"`
}
