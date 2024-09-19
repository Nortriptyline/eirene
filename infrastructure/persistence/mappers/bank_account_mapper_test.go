package mappers

import (
	"testing"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/google/uuid"
)

func TestToBankAccountDomain(t *testing.T) {
	bankID := uuid.New()

	bam := &models.BankAccountGormModel{
		ID:            uuid.New(),
		BankID:        bankID,
		OwnerID:       uuid.New().String(),
		AccountNumber: "123456789",
		AccountType:   "Checking",
		Currency:      "USD",
		Balance:       1000.0,
		Transactions:  []*models.TransactionGormModel{},
	}

	bankAccount := ToBankAccountDomain(bam)

	if bankAccount.ID != bam.ID {
		t.Errorf("Expected ID %v, got %v", bam.ID, bankAccount.ID)
	}

	if bankAccount.Bank.ID != bam.BankID {
		t.Errorf("Expected BankID %v, got %v", bam.BankID, bankAccount.Bank.ID)
	}

	if bankAccount.OwnerID != bam.OwnerID {
		t.Errorf("Expected OwnerID %v, got %v", bam.OwnerID, bankAccount.OwnerID)
	}

	if bankAccount.AccountNumber != bam.AccountNumber {
		t.Errorf("Expected AccountNumber %v, got %v", bam.AccountNumber, bankAccount.AccountNumber)
	}

	if bankAccount.AccountType != bam.AccountType {
		t.Errorf("Expected AccountType %v, got %v", bam.AccountType, bankAccount.AccountType)
	}

	if bankAccount.Currency != bam.Currency {
		t.Errorf("Expected Currency %v, got %v", bam.Currency, bankAccount.Currency)
	}

	if bankAccount.Balance != bam.Balance {
		t.Errorf("Expected Balance %v, got %v", bam.Balance, bankAccount.Balance)
	}

	// Vérification de l'absence de référence circulaire
	if bankAccount.Bank.BankAccounts != nil && len(bankAccount.Bank.BankAccounts) > 0 {
		t.Errorf("Expected Bank.ID to be %v or empty, got %v", nil, bankAccount.Bank.BankAccounts)
	}
}

func TestToBankAccountGormModel(t *testing.T) {
	bank := domain.NewBank("Test Bank", "https://testbank.com")
	ownerID := uuid.New().String()
	todayString := "2023-01-01"
	bankAccount := domain.NewBankAccount(bank, ownerID, "123456789", "Checking", "USD")
	bankAccount.CreateTransaction(100.0, domain.Deposit, "Initial deposit", todayString)

	bam := ToBankAccountGormModel(bankAccount)

	if bam.ID != bankAccount.ID {
		t.Errorf("Expected ID %v, got %v", bankAccount.ID, bam.ID)
	}

	if bam.BankID != bankAccount.Bank.ID {
		t.Errorf("Expected BankID %v, got %v", bankAccount.Bank.ID, bam.BankID)
	}

	if bam.OwnerID != bankAccount.OwnerID {
		t.Errorf("Expected OwnerID %v, got %v", bankAccount.OwnerID, bam.OwnerID)
	}

	if bam.AccountNumber != bankAccount.AccountNumber {
		t.Errorf("Expected AccountNumber %v, got %v", bankAccount.AccountNumber, bam.AccountNumber)
	}

	if bam.AccountType != bankAccount.AccountType {
		t.Errorf("Expected AccountType %v, got %v", bankAccount.AccountType, bam.AccountType)
	}

	if bam.Currency != bankAccount.Currency {
		t.Errorf("Expected Currency %v, got %v", bankAccount.Currency, bam.Currency)
	}

	if bam.Balance != bankAccount.Balance {
		t.Errorf("Expected Balance %v, got %v", bankAccount.Balance, bam.Balance)
	}

	// Vérification de l'absence de référence circulaire
	if bam.Bank.BankAccounts != nil && len(bam.Bank.BankAccounts) > 0 {
		t.Errorf("Expected Bank.ID to be %v or empty, got %v", nil, bam.Bank.BankAccounts)
	}
}

func TestToSliceBankAccountDomain(t *testing.T) {
	bankID := uuid.New()
	bams := []*models.BankAccountGormModel{
		{
			ID:            uuid.New(),
			BankID:        bankID,
			OwnerID:       uuid.New().String(),
			AccountNumber: "123456789",
			AccountType:   "Checking",
			Currency:      "USD",
			Balance:       1000.0,
			Transactions:  []*models.TransactionGormModel{},
		},
	}

	bankAccounts := ToSliceBankAccountDomain(bams)

	if len(bankAccounts) != len(bams) {
		t.Errorf("Expected %v BankAccounts, got %v", len(bams), len(bankAccounts))
	}

	for i, bankAccount := range bankAccounts {
		if bankAccount.ID != bams[i].ID {
			t.Errorf("Expected ID %v, got %v", bams[i].ID, bankAccount.ID)
		}

		if bankAccount.Bank.ID != bams[i].BankID {
			t.Errorf("Expected BankID %v, got %v", bams[i].BankID, bankAccount.Bank.ID)
		}

		if bankAccount.OwnerID != bams[i].OwnerID {
			t.Errorf("Expected OwnerID %v, got %v", bams[i].OwnerID, bankAccount.OwnerID)
		}

		if bankAccount.AccountNumber != bams[i].AccountNumber {
			t.Errorf("Expected AccountNumber %v, got %v", bams[i].AccountNumber, bankAccount.AccountNumber)
		}

		if bankAccount.AccountType != bams[i].AccountType {
			t.Errorf("Expected AccountType %v, got %v", bams[i].AccountType, bankAccount.AccountType)
		}

		if bankAccount.Currency != bams[i].Currency {
			t.Errorf("Expected Currency %v, got %v", bams[i].Currency, bankAccount.Currency)
		}

		if bankAccount.Balance != bams[i].Balance {
			t.Errorf("Expected Balance %v, got %v", bams[i].Balance, bankAccount.Balance)
		}

		// Vérification de l'absence de référence circulaire
		if bankAccount.Bank.BankAccounts != nil && len(bankAccount.Bank.BankAccounts) > 0 {
			t.Errorf("Expected Bank.ID to be %v or empty, got %v", nil, bankAccount.Bank.BankAccounts)
		}
	}
}

func TestToSliceBankAccountGormModel(t *testing.T) {
	bank := domain.NewBank("Test Bank", "https://testbank.com")
	ownerID := uuid.New().String()
	todayString := "2023-01-01"
	bankAccount := domain.NewBankAccount(bank, ownerID, "123456789", "Checking", "USD")
	bankAccount.CreateTransaction(100.0, domain.Deposit, "Initial deposit", todayString)

	bankAccounts := []*domain.BankAccount{bankAccount}

	bams := ToSliceBankAccountGormModel(bankAccounts)

	if len(bams) != len(bankAccounts) {
		t.Errorf("Expected %v BankAccountGormModels, got %v", len(bankAccounts), len(bams))
	}

	for i, bam := range bams {
		if bam.ID != bankAccounts[i].ID {
			t.Errorf("Expected ID %v, got %v", bankAccounts[i].ID, bam.ID)
		}

		if bam.BankID != bankAccounts[i].Bank.ID {
			t.Errorf("Expected BankID %v, got %v", bankAccounts[i].Bank.ID, bam.BankID)
		}

		if bam.OwnerID != bankAccounts[i].OwnerID {
			t.Errorf("Expected OwnerID %v, got %v", bankAccounts[i].OwnerID, bam.OwnerID)
		}

		if bam.AccountNumber != bankAccounts[i].AccountNumber {
			t.Errorf("Expected AccountNumber %v, got %v", bankAccounts[i].AccountNumber, bam.AccountNumber)
		}

		if bam.AccountType != bankAccounts[i].AccountType {
			t.Errorf("Expected AccountType %v, got %v", bankAccounts[i].AccountType, bam.AccountType)
		}

		if bam.Currency != bankAccounts[i].Currency {
			t.Errorf("Expected Currency %v, got %v", bankAccounts[i].Currency, bam.Currency)
		}

		if bam.Balance != bankAccounts[i].Balance {
			t.Errorf("Expected Balance %v, got %v", bankAccounts[i].Balance, bam.Balance)
		}

		// Vérification de l'absence de référence circulaire
		if bam.Bank.BankAccounts != nil && len(bam.Bank.BankAccounts) > 0 {
			t.Errorf("Expected Bank.ID to be %v or empty, got %v", nil, bam.Bank.BankAccounts)
		}
	}
}
