package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"` // Primary key
	ClientToken string    `gorm:"not null"`             // Auto Generated token by the client (browser; app; etc)
	CreatedAt   time.Time `gorm:"autoCreateTime"`       // Creation timestamp
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users"
}

func (r *UserModel) BeforeCreate(*gorm.DB) error {
	r.ID = uuid.New()

	return nil
}
