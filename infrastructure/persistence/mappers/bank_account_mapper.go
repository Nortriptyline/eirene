package mappers

import (
	"fmt"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
)

func ToBankAccountDomain(bam *models.BankAccountGormModel) *domain.BankAccount {
	// Implement the mapping logic from BankAccountGormModel to BankAccount
	return &domain.BankAccount{
		ID:            bam.ID,
		Bank:          &domain.Bank{ID: bam.BankID}, // Assuming Bank is a field in BankAccountGormModel
		OwnerID:       bam.OwnerID,
		AccountNumber: bam.AccountNumber,
		AccountType:   bam.AccountType,
		Currency:      bam.Currency,
		Balance:       bam.Balance,
		Transactions:  ToSliceTransactionDomain(bam.Transactions), // You may want to map transactions as well
	}
}

func ToSliceBankAccountDomain(bams []*models.BankAccountGormModel) []*domain.BankAccount {
	bankAccounts := make([]*domain.BankAccount, len(bams))
	for i, bam := range bams {
		bankAccounts[i] = ToBankAccountDomain(bam)
	}
	return bankAccounts
}

func ToBankAccountGormModel(bankAccount *domain.BankAccount) *models.BankAccountGormModel {
	// Implement the mapping logic from BankAccount to BankAccountGormModel
	return &models.BankAccountGormModel{
		ID:            bankAccount.ID,
		BankID:        bankAccount.Bank.ID, // Assuming Bank has an ID field
		OwnerID:       bankAccount.OwnerID,
		AccountNumber: bankAccount.AccountNumber,
		AccountType:   bankAccount.AccountType,
		Currency:      bankAccount.Currency,
		Balance:       bankAccount.Balance,
		Transactions:  nil, // You may want to map transactions as well
	}
}

func ToSliceBankAccountGormModel(bankAccounts []*domain.BankAccount) []*models.BankAccountGormModel {
	bankAccountModels := make([]*models.BankAccountGormModel, len(bankAccounts))
	for i, bankAccount := range bankAccounts {
		fmt.Print(bankAccount.Bank)
		bankAccountModels[i] = ToBankAccountGormModel(bankAccount)
	}
	return bankAccountModels
}
