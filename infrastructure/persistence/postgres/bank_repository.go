package postgres

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/mappers"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankRepositoryImpl struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) repository.IBankRepository {
	return &BankRepositoryImpl{db: db}
}

func (r *BankRepositoryImpl) Create(bank *domain.Bank) error {
	// Map domain.Bank to the GORM model if necessary
	b := mappers.ToBankGormModel(bank)

	return r.db.Create(b).Error
}

func (r *BankRepositoryImpl) FindByID(id uuid.UUID) (*domain.Bank, error) {
	var bankGormModel models.BankGormModel
	if err := r.db.First(&bankGormModel, "id = ?", id).Error; err != nil {
		return nil, err
	}

	bank := mappers.ToBankDomain(&bankGormModel)

	return bank, nil
}

func (r *BankRepositoryImpl) Update(bank *domain.Bank) error {
	b := mappers.ToBankGormModel(bank)
	return r.db.Save(b).Error
}

func (r *BankRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Bank{}, "id = ?", id).Error
}
