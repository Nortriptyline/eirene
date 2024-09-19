package postgres

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/mappers"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models/model_errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repository.ITransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) Create(transaction *domain.Transaction) error {
	// Map domain.Transaction to the GORM model if necessary
	t := mappers.ToTransactionGormModel(transaction)
	return r.db.Create(&t).Error
}

func (r *TransactionRepositoryImpl) Update(transaction *domain.Transaction) error {
	t := mappers.ToTransactionGormModel(transaction)
	return r.db.Save(&t).Error
}

func (r *TransactionRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.TransactionGormModel{}, "id = ?", id).Error
}

func (r *TransactionRepositoryImpl) FindByID(id uuid.UUID) (*domain.Transaction, error) {
	var transaction models.TransactionGormModel
	if err := r.db.First(&transaction, "id = ?", id).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, model_errors.ErrDbTransactionNotFound
		}

		return nil, err
	}

	return mappers.ToTransactionDomain(&transaction), nil
}

func (r *TransactionRepositoryImpl) FindAll() ([]*domain.Transaction, error) {
	var transactions []*models.TransactionGormModel
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return mappers.ToSliceTransactionDomain(transactions), nil
}
