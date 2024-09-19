package repository

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/google/uuid"
)

type IBankRepository interface {
	Create(bank *domain.Bank) error
	FindByID(id uuid.UUID) (*domain.Bank, error)
	Update(bank *domain.Bank) error
	Delete(id uuid.UUID) error
}
