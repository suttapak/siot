package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Box struct {
	ID        uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `sql:"index" json:"deletedAt"`
	// attribute
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     uuid.UUID `json:"ownerId"`

	// relation
	Members   []BoxMember  `gorm:"foreignKey:BoxId" json:"members"`
	BoxSecret BoxSecret    `gorm:"foreignKey:BoxId" json:"boxSecret"`
	CanSub    CanSubscribe `gorm:"foreignKey:BoxId" json:"canSub"`
	CanPub    CanPublish   `gorm:"foreignKey:BoxId" json:"canPub"`
}

func (b *Box) BeforeCreate(*gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
