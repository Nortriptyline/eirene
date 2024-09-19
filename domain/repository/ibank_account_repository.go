package repository

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/google/uuid"
)

// BankAccountRepository définit les méthodes pour interagir avec les comptes bancaires.

type IBankAccountRepository interface {
	// Create ajoute un nouveau compte bancaire.
	Create(bankAccount *domain.BankAccount) error

	// Update met à jour un compte bancaire existant.
	Update(bankAccount *domain.BankAccount) error

	// Delete supprime un compte bancaire par son ID.
	Delete(id uuid.UUID) error

	// FindByID recherche un compte bancaire par son ID.
	FindByID(id uuid.UUID) (*domain.BankAccount, error)

	// FindAll retourne tous les comptes bancaires.
	FindAll() ([]*domain.BankAccount, error)
}
