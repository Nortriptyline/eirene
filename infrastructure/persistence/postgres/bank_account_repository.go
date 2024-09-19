package postgres

import (
	"errors"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/mappers"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models/model_errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccountRepositoryImpl struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) repository.IBankAccountRepository {
	return &BankAccountRepositoryImpl{db: db}
}

func (r *BankAccountRepositoryImpl) Create(bankAccount *domain.BankAccount) error {
	// Map domain.BankAccount to the GORM model if necessary
	model := mappers.ToBankAccountGormModel(bankAccount)
	err := r.db.Create(&model).Error

	// If the error is a duplicated key error, it means that the bank account already exists
	// This should not happen because the domain layer should prevent this from happening
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return model_errors.ErrDbBankAccountAlreadyExists
	}

	return err
}

func (r *BankAccountRepositoryImpl) Update(bankAccount *domain.BankAccount) error {
	model := mappers.ToBankAccountGormModel(bankAccount)
	return r.db.Save(&model).Error
}

func (r *BankAccountRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.BankAccountGormModel{}, "id = ?", id).Error
}

func (r *BankAccountRepositoryImpl) FindByID(id uuid.UUID) (*domain.BankAccount, error) {
	var model models.BankAccountGormModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model_errors.ErrDbBankAccountNotFound
		}

		return nil, err
	}

	bankAccount := mappers.ToBankAccountDomain(&model)

	return bankAccount, nil
}

func (r *BankAccountRepositoryImpl) FindAll() ([]*domain.BankAccount, error) {
	var m []*models.BankAccountGormModel
	if err := r.db.Find(&m).Error; err != nil {
		return nil, err
	}

	bankAccounts := mappers.ToSliceBankAccountDomain(m)

	return bankAccounts, nil
}
