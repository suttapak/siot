package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// common
	ID        uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `sql:"index" json:"deletedAt"`
	// attribute
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"index:unique" json:"email"`
	Password  string `json:"password"`
	Avatar    Avatar `json:"avatar" gorm:"foreignKey:UserId"`
	// fk
	SettingId uint `json:"settingId"`
	// relation
	Setting UserSetting `gorm:"foreignKey:UserId" json:"setting"`
	Roles   []*Role     `gorm:"many2many:user_roles;" json:"roles"`
	Box     []*Box      `gorm:"foreignKey:OwnerId" json:"box"`
}

func (m *User) BeforeCreate(*gorm.DB) error {
	m.ID = uuid.New()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(m.Password), viper.GetInt("pw.hash.salt"))
	if err != nil {
		return err
	}
	m.Password = string(passwordHash)
	return nil
}

func (m *User) PasswordIsCorrect(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password)) == nil
}

func (m *User) EmailIsExist() bool {
	// ถ้ามีรีเทิน ture
	return len(m.Email) > 0
}
