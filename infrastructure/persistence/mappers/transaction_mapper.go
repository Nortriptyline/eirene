package mappers

import (
	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
)

func ToTransactionDomain(tm *models.TransactionGormModel) *domain.Transaction {
	// Implement the mapping logic from TransactionGormModel to Transaction
	return &domain.Transaction{
		ID:              tm.ID,
		BankAccount:     &domain.BankAccount{ID: tm.BankAccountID}, // Assuming BankAccountID is a field in TransactionGormModel
		Amount:          tm.Amount,
		Description:     tm.Description,
		TransactionType: tm.TransactionType,
		TransactionDate: tm.TransactionDate,
	}
}

func ToSliceTransactionDomain(tms []*models.TransactionGormModel) []*domain.Transaction {
	transactions := make([]*domain.Transaction, len(tms))
	for i, tm := range tms {
		transactions[i] = ToTransactionDomain(tm)
	}
	return transactions
}

func ToTransactionGormModel(transaction *domain.Transaction) *models.TransactionGormModel {
	// Implement the mapping logic from Transaction to TransactionGormModel
	return &models.TransactionGormModel{
		ID:              transaction.ID,
		BankAccountID:   transaction.BankAccount.ID, // Assuming BankAccount has an ID field
		Amount:          transaction.Amount,
		Description:     transaction.Description,
		TransactionType: transaction.TransactionType,
		TransactionDate: transaction.TransactionDate,
	}
}

func ToSliceTransactionGormModel(transactions []*domain.Transaction) []*models.TransactionGormModel {
	transactionModels := make([]*models.TransactionGormModel, len(transactions))
	for i, transaction := range transactions {
		transactionModels[i] = ToTransactionGormModel(transaction)
	}
	return transactionModels
}
