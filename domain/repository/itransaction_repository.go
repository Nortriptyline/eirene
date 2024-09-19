package repository

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/google/uuid"
)

type ITransactionRepository interface {
	// Create ajoute une nouvelle transaction.
	Create(transaction *domain.Transaction) error

	// Update met à jour une transaction existante.
	Update(transaction *domain.Transaction) error

	// Delete supprime une transaction par son ID.
	Delete(id uuid.UUID) error

	// FindByID recherche une transaction par son ID.
	FindByID(id uuid.UUID) (*domain.Transaction, error)

	// FindAll retourne toutes les transactions.
	FindAll() ([]*domain.Transaction, error)
}
