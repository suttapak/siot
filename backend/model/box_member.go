package model

import "github.com/google/uuid"

type BoxMember struct {
	Model
	// attribute
	UserAccessToken string `json:"userAccessToken"`
	// fk
	UserId uuid.UUID `json:"userId"`
	User   User      `json:"user"`
	BoxId  uuid.UUID `json:"boxId"`
	// relation
	BoxMemberPermission *BoxMemberPermission `gorm:"foreignKey:BoxMemberId" json:"boxMemberPermission"`
}
