package mappers

import (
	"testing"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/google/uuid"
)

func TestToBankDomain(t *testing.T) {
	bankAccountID := uuid.New()
	bgm := &models.BankGormModel{
		ID:      uuid.New(),
		Name:    "Test Bank",
		Website: "https://testbank.com",
		BankAccounts: []*models.BankAccountGormModel{
			{
				ID:            bankAccountID,
				AccountNumber: "123456789",
				AccountType:   "Checking",
				Currency:      "USD",
			},
		},
	}

	bank := ToBankDomain(bgm)

	if bank.ID != bgm.ID {
		t.Errorf("Expected ID %v, got %v", bgm.ID, bank.ID)
	}

	if bank.Name != bgm.Name {
		t.Errorf("Expected Name %v, got %v", bgm.Name, bank.Name)
	}

	if bank.Website != bgm.Website {
		t.Errorf("Expected Website %v, got %v", bgm.Website, bank.Website)
	}

	if len(bank.BankAccounts) != len(bgm.BankAccounts) {
		t.Errorf("Expected %v BankAccounts, got %v", len(bgm.BankAccounts), len(bank.BankAccounts))
	}

	if bank.BankAccounts[0].ID != bgm.BankAccounts[0].ID {
		t.Errorf("Expected BankAccount ID %v, got %v", bgm.BankAccounts[0].ID, bank.BankAccounts[0].ID)
	}
}

func TestToBankGormModel(t *testing.T) {

	bankID := uuid.New()

	bank := domain.NewBank("Test Bank", "https://testbank.com")
	bank.ID = bankID

	account := domain.NewBankAccount(bank, "123", "123", "checking", "USD")
	account.ID = uuid.New()
	bank.BankAccounts = append(bank.BankAccounts, account)

	bgm := ToBankGormModel(bank)

	if bgm.ID != bank.ID {
		t.Errorf("Expected ID %v, got %v", bank.ID, bgm.ID)
	}

	if bgm.Name != bank.Name {
		t.Errorf("Expected Name %v, got %v", bank.Name, bgm.Name)
	}

	if bgm.Website != bank.Website {
		t.Errorf("Expected Website %v, got %v", bank.Website, bgm.Website)
	}

	if len(bgm.BankAccounts) != len(bank.BankAccounts) {
		t.Errorf("Expected %v BankAccounts, got %v", len(bank.BankAccounts), len(bgm.BankAccounts))
	}

	if bgm.BankAccounts[0].ID != bank.BankAccounts[0].ID {
		t.Errorf("Expected BankAccount ID %v, got %v", bank.BankAccounts[0].ID, bgm.BankAccounts[0].ID)
	}

	// Make sure their is no circular reference
	if bgm.BankAccounts[0].Bank.BankAccounts != nil {
		t.Errorf("Expected BankAccounts to be nil, got %v", bgm.BankAccounts[0].Bank.BankAccounts)
	}
}
