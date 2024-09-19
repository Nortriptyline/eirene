package application

import (
	"github.com/Nortriptyline/Eirene/domain/repository"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	BankRepository        repository.IBankRepository
	BankAccountRepository repository.IBankAccountRepository
	TransactionRepository repository.ITransactionRepository
}

func InitializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		BankRepository:        postgres.NewBankRepository(db),
		BankAccountRepository: postgres.NewBankAccountRepository(db),
		TransactionRepository: postgres.NewTransactionRepository(db),
	}
}
