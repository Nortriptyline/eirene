package mappers

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
)

func ToBankDomain(bgm *models.BankGormModel) *domain.Bank {
	return &domain.Bank{
		ID:           bgm.ID,
		Name:         bgm.Name,
		Website:      bgm.Website,
		BankAccounts: ToSliceBankAccountDomain(bgm.BankAccounts), // Assuming BankAccounts is a field in BankGormModel
	}
}

func ToBankGormModel(b *domain.Bank) *models.BankGormModel {
	return &models.BankGormModel{
		ID:           b.ID,
		Name:         b.Name,
		Website:      b.Website,
		BankAccounts: ToSliceBankAccountGormModel(b.BankAccounts), // Assuming BankAccounts is a field in Bank
	}
}
