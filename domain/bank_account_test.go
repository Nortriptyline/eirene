package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewBankAccount(t *testing.T) {
	bank := &Bank{
		ID:      uuid.New(),
		Name:    "Test Bank",
		Website: "https://testbank.com",
	}

	ownerID := "owner123"
	accountNumber := "123456789"
	accountType := Checking
	currency := "USD"

	bankAccount := NewBankAccount(bank, ownerID, accountNumber, accountType, currency)

	if bankAccount.ID == uuid.Nil {
		t.Errorf("Expected a valid UUID, got %v", bankAccount.ID)
	}

	if bankAccount.Bank != bank {
		t.Errorf("Expected bank %v, got %v", bank, bankAccount.Bank)
	}

	if bankAccount.OwnerID != ownerID {
		t.Errorf("Expected owner ID %v, got %v", ownerID, bankAccount.OwnerID)
	}

	if bankAccount.AccountNumber != accountNumber {
		t.Errorf("Expected account number %v, got %v", accountNumber, bankAccount.AccountNumber)
	}

	if bankAccount.AccountType != accountType {
		t.Errorf("Expected account type %v, got %v", accountType, bankAccount.AccountType)
	}

	if len(bankAccount.Transactions) != 0 {
		t.Errorf("Expected empty transactions, got %v", len(bankAccount.Transactions))
	}

	if bankAccount.Balance != 0.0 {
		t.Errorf("Expected balance 0.0, got %v", bankAccount.Balance)
	}

	if bankAccount.Currency != currency {
		t.Errorf("Expected currency %v, got %v", currency, bankAccount.Currency)
	}
}

func TestAddTransaction(t *testing.T) {
	bank := NewBank("Test Bank", "https://testbank.com")
	bankAccount := NewBankAccount(bank, "owner123", "123456789", Checking, "USD")

	transaction := &Transaction{
		ID:              uuid.New(),
		BankAccount:     bankAccount,
		Amount:          100.0,
		TransactionType: Deposit,
		Description:     "Test deposit",
		TransactionDate: "2023-10-01",
	}

	bankAccount.AddTransaction(transaction)

	if len(bankAccount.Transactions) != 1 {
		t.Errorf("Expected 1 transaction, got %v", len(bankAccount.Transactions))
	}

	assert.Equal(t, transaction.ID, bankAccount.Transactions[0].ID)
	assert.Equal(t, transaction.BankAccount.ID, bankAccount.Transactions[0].BankAccount.ID)
	assert.Equal(t, transaction.Amount, bankAccount.Transactions[0].Amount)
	assert.Equal(t, transaction.TransactionType, bankAccount.Transactions[0].TransactionType)
	assert.Equal(t, transaction.Description, bankAccount.Transactions[0].Description)
	assert.Equal(t, transaction.TransactionDate, bankAccount.Transactions[0].TransactionDate)
}

func TestCreateTransaction(t *testing.T) {
	bank := NewBank("Test Bank", "https://testbank.com")
	bankAccount := NewBankAccount(bank, "owner123", "123456789", Checking, "USD")

	amount := 100.0
	transactionType := Deposit
	description := "Test deposit"
	date := "2023-10-01"

	transaction := bankAccount.CreateTransaction(amount, transactionType, description, date)

	assert.NotNil(t, transaction)
	assert.NotEqual(t, transaction.ID, uuid.Nil)
	assert.Equal(t, transaction.BankAccount, bankAccount)
	assert.Equal(t, transaction.Amount, amount)
	assert.Equal(t, transaction.TransactionType, transactionType)
	assert.Equal(t, transaction.Description, description)
	assert.Equal(t, transaction.TransactionDate, date)
	assert.Len(t, bankAccount.Transactions, 1)
	assert.Equal(t, transaction, bankAccount.Transactions[0])
}
