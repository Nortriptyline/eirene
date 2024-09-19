package mappers

import (
	"testing"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/google/uuid"
)

func TestToTransactionDomain(t *testing.T) {
	bankAccountID := uuid.New()
	tm := &models.TransactionGormModel{
		ID:              uuid.New(),
		BankAccountID:   bankAccountID,
		Amount:          100.0,
		Description:     "Test transaction",
		TransactionType: "Credit",
	}

	transaction := ToTransactionDomain(tm)

	if transaction.ID != tm.ID {
		t.Errorf("Expected ID %v, got %v", tm.ID, transaction.ID)
	}

	if transaction.BankAccount.ID != tm.BankAccountID {
		t.Errorf("Expected BankAccountID %v, got %v", tm.BankAccountID, transaction.BankAccount.ID)
	}

	if transaction.Amount != tm.Amount {
		t.Errorf("Expected Amount %v, got %v", tm.Amount, transaction.Amount)
	}

	if transaction.Description != tm.Description {
		t.Errorf("Expected Description %v, got %v", tm.Description, transaction.Description)
	}

	if transaction.TransactionType != tm.TransactionType {
		t.Errorf("Expected TransactionType %v, got %v", tm.TransactionType, transaction.TransactionType)
	}

	if transaction.BankAccount.Transactions != nil && len(transaction.BankAccount.Transactions) > 0 {
		t.Errorf("Expected BankAccount.Transactions to be %v or empty, got %v", nil, transaction.BankAccount.Transactions)
	}
}

func TestToSliceTransactionDomain(t *testing.T) {
	bankAccountID := uuid.New()
	tms := []*models.TransactionGormModel{
		{
			ID:              uuid.New(),
			BankAccountID:   bankAccountID,
			Amount:          100.0,
			Description:     "Test transaction 1",
			TransactionType: "Credit",
		},
		{
			ID:              uuid.New(),
			BankAccountID:   bankAccountID,
			Amount:          200.0,
			Description:     "Test transaction 2",
			TransactionType: "Debit",
		},
	}

	transactions := ToSliceTransactionDomain(tms)

	if len(transactions) != len(tms) {
		t.Errorf("Expected %v transactions, got %v", len(tms), len(transactions))
	}

	for i, transaction := range transactions {
		if transaction.ID != tms[i].ID {
			t.Errorf("Expected ID %v, got %v", tms[i].ID, transaction.ID)
		}

		if transaction.BankAccount.ID != tms[i].BankAccountID {
			t.Errorf("Expected BankAccountID %v, got %v", tms[i].BankAccountID, transaction.BankAccount.ID)
		}

		if transaction.Amount != tms[i].Amount {
			t.Errorf("Expected Amount %v, got %v", tms[i].Amount, transaction.Amount)
		}

		if transaction.Description != tms[i].Description {
			t.Errorf("Expected Description %v, got %v", tms[i].Description, transaction.Description)
		}

		if transaction.TransactionType != tms[i].TransactionType {
			t.Errorf("Expected TransactionType %v, got %v", tms[i].TransactionType, transaction.TransactionType)
		}

		if transaction.BankAccount.Transactions != nil && len(transaction.BankAccount.Transactions) > 0 {
			t.Errorf("Expected BankAccount.Transactions to be %v or empty, got %v", nil, transaction.BankAccount.Transactions)
		}
	}
}

func TestToTransactionGormModel(t *testing.T) {
	bank := domain.NewBank("Test Bank", "https://testbank.com")
	bankAccount := domain.NewBankAccount(bank, uuid.New().String(), "123456789", "Checking", "USD")

	bankAccount.CreateTransaction(100.0, domain.Deposit, "Initial deposit", "2023-01-01")
	transaction := bankAccount.Transactions[0]

	tm := ToTransactionGormModel(transaction)

	if tm.ID != transaction.ID {
		t.Errorf("Expected ID %v, got %v", transaction.ID, tm.ID)
	}

	if tm.BankAccountID != transaction.BankAccount.ID {
		t.Errorf("Expected BankAccountID %v, got %v", transaction.BankAccount.ID, tm.BankAccountID)
	}

	if tm.Amount != transaction.Amount {
		t.Errorf("Expected Amount %v, got %v", transaction.Amount, tm.Amount)
	}

	if tm.Description != transaction.Description {
		t.Errorf("Expected Description %v, got %v", transaction.Description, tm.Description)
	}

	if tm.TransactionType != transaction.TransactionType {
		t.Errorf("Expected TransactionType %v, got %v", transaction.TransactionType, tm.TransactionType)
	}

	if tm.BankAccount.Transactions != nil && len(tm.BankAccount.Transactions) > 0 {
		t.Errorf("Expected BankAccount.Transactions to be %v or empty, got %v", nil, tm.BankAccount.Transactions)
	}
}

func TestToSliceTransactionGormModel(t *testing.T) {
	bank := domain.NewBank("Test Bank", "https://testbank.com")
	bankAccount := domain.NewBankAccount(bank, uuid.New().String(), "123456789", "Checking", "USD")

	bankAccount.CreateTransaction(100.0, domain.Deposit, "Initial deposit", "2023-01-01")
	transactions := bankAccount.Transactions

	tms := ToSliceTransactionGormModel(transactions)

	if len(tms) != len(transactions) {
		t.Errorf("Expected %v transaction models, got %v", len(transactions), len(tms))
	}

	for i, tm := range tms {
		if tm.ID != transactions[i].ID {
			t.Errorf("Expected ID %v, got %v", transactions[i].ID, tm.ID)
		}

		if tm.BankAccountID != transactions[i].BankAccount.ID {
			t.Errorf("Expected BankAccountID %v, got %v", transactions[i].BankAccount.ID, tm.BankAccountID)
		}

		if tm.Amount != transactions[i].Amount {
			t.Errorf("Expected Amount %v, got %v", transactions[i].Amount, tm.Amount)
		}

		if tm.Description != transactions[i].Description {
			t.Errorf("Expected Description %v, got %v", transactions[i].Description, tm.Description)
		}

		if tm.TransactionType != transactions[i].TransactionType {
			t.Errorf("Expected TransactionType %v, got %v", transactions[i].TransactionType, tm.TransactionType)
		}

		if tm.BankAccount.Transactions != nil && len(tm.BankAccount.Transactions) > 0 {
			t.Errorf("Expected BankAccount.Transactions to be %v or empty, got %v", nil, tm.BankAccount.Transactions)
		}
	}
}
