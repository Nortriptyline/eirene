package models

import (
	"time"

	"github.com/google/uuid"
)

type BankGormModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name         string    `gorm:"not null"`
	Website      string
	BankAccounts []*BankAccountGormModel `gorm:"foreignKey:BankID"`
	CreatedAt    time.Time               `gorm:"autoCreateTime"`
	UpdatedAt    time.Time               `gorm:"autoUpdateTime"`
}

func (BankGormModel) TableName() string {
	return "banks"
}
