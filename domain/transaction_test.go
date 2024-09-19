// transaction_test.go
package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTransaction(t *testing.T) {
	bank := NewBank("Test Bank", "https://testbank.com")

	bankAccount := &BankAccount{
		ID:            uuid.New(),
		Bank:          bank,
		OwnerID:       "owner123",
		AccountNumber: "123456789",
		AccountType:   Checking,
		Currency:      "USD",
	}

	amount := 100.0
	transactionType := Deposit
	description := "Test deposit"
	transactionDate := "2023-10-01"

	transaction := NewTransaction(bankAccount, amount, transactionType, description, transactionDate)

	if transaction.ID == uuid.Nil {
		t.Errorf("Expected a valid UUID, got %v", transaction.ID)
	}

	if transaction.BankAccount.ID != bankAccount.ID {
		t.Errorf("Expected bank account %v, got %v", bankAccount, transaction.BankAccount)
	}

	if transaction.Amount != amount {
		t.Errorf("Expected amount %v, got %v", amount, transaction.Amount)
	}

	if transaction.TransactionType != transactionType {
		t.Errorf("Expected transaction type %v, got %v", transactionType, transaction.TransactionType)
	}

	if transaction.Description != description {
		t.Errorf("Expected description %v, got %v", description, transaction.Description)
	}

	if transaction.TransactionDate != transactionDate {
		t.Errorf("Expected transaction date %v, got %v", transactionDate, transaction.TransactionDate)
	}
}
