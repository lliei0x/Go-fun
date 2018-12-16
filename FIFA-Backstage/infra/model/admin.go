package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Name              string `gorm:"type:varchar;column:name"`
	AuthToken         string `gorm:"type:varchar;column:auth_token"`
	EncryptedPassword string `gorm:"type:varchar;column:encrypted_password'"`
	Phone             string `gorm:"type:varchar;column:phone"`
	State             string `gorm:"type:varchar;column:state"`
}

type AdminSerializer struct {
	ID                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Name              string     `json:"name"`
	AuthToken         string     `json:"auth_token"`
	EncryptedPassword string     `json:"encrypted_password"`
	Phone             string     `json:"phone"`
	State             string     `json:"state"`
}

func (a *Admin) Serializer() AdminSerializer {
	return AdminSerializer{
		ID:                a.ID,
		CreatedAt:         a.CreatedAt,
		UpdatedAt:         a.UpdatedAt,
		DeletedAt:         a.DeletedAt,
		Name:              a.Name,
		AuthToken:         a.AuthToken,
		EncryptedPassword: a.EncryptedPassword,
		Phone:             a.Phone,
		State:             a.State,
	}
}
