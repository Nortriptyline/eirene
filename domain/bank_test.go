// bank_test.go
package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewBank(t *testing.T) {
	name := "Test Bank"
	website := "https://testbank.com"
	bank := NewBank(name, website)

	if bank.ID == uuid.Nil {
		t.Errorf("Expected a valid UUID, got %v", bank.ID)
	}

	if bank.Name != name {
		t.Errorf("Expected name %v, got %v", name, bank.Name)
	}

	if bank.Website != website {
		t.Errorf("Expected website %v, got %v", website, bank.Website)
	}

	if len(bank.BankAccounts) != 0 {
		t.Errorf("Expected empty bank accounts, got %v", len(bank.BankAccounts))
	}
}

func TestAddNewBankAccount(t *testing.T) {
	bank := NewBank("Test Bank", "https://testbank.com")
	ownerID := "owner123"
	accountNumber := "123456789"
	accountType := Checking
	currency := "USD"

	bankAccount := bank.AddNewBankAccount(ownerID, accountNumber, accountType, currency)

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

	if len(bank.BankAccounts) != 1 {
		t.Errorf("Expected 1 bank account in bank, got %v", len(bank.BankAccounts))
	}

	if bank.BankAccounts[0] != bankAccount {
		t.Errorf("Expected bank account %v in bank, got %v", bankAccount, bank.BankAccounts[0])
	}
}

func TestAddBankAccount(t *testing.T) {
	bank := NewBank("Test Bank", "https://testbank.com")
	bankAccount := NewBankAccount(bank, "owner123", "123456789", Checking, "USD")

	bank.AddBankAccount(bankAccount)

	if len(bank.BankAccounts) != 1 {
		t.Errorf("Expected 1 bank account, got %v", len(bank.BankAccounts))
	}

	if bank.BankAccounts[0] != bankAccount {
		t.Errorf("Expected bank account %v, got %v", bankAccount, bank.BankAccounts[0])
	}
}
