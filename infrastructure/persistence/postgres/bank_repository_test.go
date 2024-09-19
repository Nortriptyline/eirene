package postgres

import (
	"testing"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/mappers"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/google/uuid"
)

func TestCreateBank(t *testing.T) {
	db := SetupTestDB()
	repo := NewBankRepository(db)

	bank := &domain.Bank{
		ID:      uuid.New(),
		Name:    "Test Bank",
		Website: "https://testbank.com",
	}

	err := repo.Create(bank)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var bankGormModel models.BankGormModel
	if err := db.First(&bankGormModel, "id = ?", bank.ID).Error; err != nil {
		t.Fatalf("Expected to find bank, got error %v", err)
	}

	if bankGormModel.Name != bank.Name {
		t.Errorf("Expected name %v, got %v", bank.Name, bankGormModel.Name)
	}

	if bankGormModel.Website != bank.Website {
		t.Errorf("Expected website %v, got %v", bank.Website, bankGormModel.Website)
	}
}

func TestFindByID(t *testing.T) {
	db := SetupTestDB()
	repo := NewBankRepository(db)

	bank := &domain.Bank{
		ID:      uuid.New(),
		Name:    "Test Bank",
		Website: "https://testbank.com",
	}

	bankGormModel := mappers.ToBankGormModel(bank)
	db.Create(bankGormModel)

	foundBank, err := repo.FindByID(bank.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if foundBank.Name != bank.Name {
		t.Errorf("Expected name %v, got %v", bank.Name, foundBank.Name)
	}

	if foundBank.Website != bank.Website {
		t.Errorf("Expected website %v, got %v", bank.Website, foundBank.Website)
	}
}
