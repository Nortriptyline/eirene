package models

import (
	"time"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models/model_errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccountGormModel struct {
	ID            uuid.UUID               `gorm:"type:uuid;primaryKey"`
	BankID        uuid.UUID               `gorm:"type:uuid;not null"`
	Bank          *BankGormModel          `gorm:"foreignKey:BankID"`
	OwnerID       string                  `gorm:"not null;"`
	AccountNumber string                  `gorm:"not null;unique"`
	AccountType   domain.AccountType      `gorm:"not null"`
	Currency      string                  `gorm:"not null"`
	Balance       float64                 `gorm:"default:0"`
	Transactions  []*TransactionGormModel `gorm:"foreignKey:BankAccountID"`
	CreatedAt     time.Time               `gorm:"autoCreateTime"`
	UpdatedAt     time.Time               `gorm:"autoUpdateTime"`
}

func (BankAccountGormModel) TableName() string {
	return "bank_accounts"
}

// BeforeCreate est un hook GORM qui est appelé avant la création d'un enregistrement
func (b *BankAccountGormModel) BeforeCreate(tx *gorm.DB) error {
	if err := b.validate(); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate est un hook GORM qui est appelé avant la mise à jour d'un enregistrement
func (b *BankAccountGormModel) BeforeUpdate(tx *gorm.DB) error {
	if err := b.validate(); err != nil {
		return err
	}
	return nil
}

// validate contient toutes les validations pour le modèle BankAccountGormModel
func (b *BankAccountGormModel) validate() error {
	if b.OwnerID == "" {
		return model_errors.ErrDbBankAccountOwnedIDEmpty
	}
	if b.AccountNumber == "" {
		return model_errors.ErrDbBankAccountAccountNumberEmpty
	}
	if b.Currency == "" {
		return model_errors.ErrDbBankAccountCurrencyInvalid
	}

	// Verify type is valid
	if !b.AccountType.IsValid() {
		return model_errors.ErrDbBankAccountInvalidType
	}

	return nil
}
