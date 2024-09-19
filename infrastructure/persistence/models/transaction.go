package models

import (
	"time"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/google/uuid"
)

type TransactionGormModel struct {
	ID              uuid.UUID             `gorm:"type:uuid;primaryKey"`
	BankAccountID   uuid.UUID             `gorm:"type:uuid;not null"`
	BankAccount     *BankAccountGormModel `gorm:"foreignKey:BankAccountID"`
	Amount          float64               `gorm:"not null"`
	Description     string
	TransactionType domain.TransactionType `gorm:"not null"`
	TransactionDate string                 `gorm:"not null"`
	CreatedAt       time.Time              `gorm:"autoCreateTime"`
	UpdatedAt       time.Time              `gorm:"autoUpdateTime"`
}

func (TransactionGormModel) TableName() string {
	return "transactions"
}
