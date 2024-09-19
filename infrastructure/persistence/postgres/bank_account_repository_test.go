package postgres

import (
	"testing"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models/model_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BankAccountRepositoryTest struct {
	repoBank repository.IBankRepository
	repoBA   repository.IBankAccountRepository
	bank     *domain.Bank
	db       *gorm.DB
}

func NewBankAccountRepositoryTest() *BankAccountRepositoryTest {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("PRAGMA foreign_keys = ON")

	// Auto migrate the schema
	db.AutoMigrate(
		&models.BankGormModel{},
		&models.BankAccountGormModel{},
	)

	bank := domain.NewBank("Test Bank", "https://testbank.com")
	repoBank := NewBankRepository(db)
	repoBA := NewBankAccountRepository(db)

	err = repoBank.Create(bank)
	if err != nil {
		panic(err)
	}

	return &BankAccountRepositoryTest{
		repoBank: repoBank,
		repoBA:   repoBA,
		bank:     bank,
		db:       db,
	}
}

func TestCreateBankAccount(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Create a bank account
	ownerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bat.bank, ownerID, "account_number", domain.Checking, "USD")

	// Save the bank account (function to be tested)
	err := bat.repoBA.Create(bankAccount)
	assert.NoError(t, err)

	// Retrieve the bank account from the database
	var bankAccountFromDB models.BankAccountGormModel
	err = bat.db.First(&bankAccountFromDB, "id = ?", bankAccount.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, bankAccount.ID, bankAccountFromDB.ID)
	assert.Equal(t, bankAccount.OwnerID, bankAccountFromDB.OwnerID)
	assert.Equal(t, bankAccount.AccountNumber, bankAccountFromDB.AccountNumber)
	assert.Equal(t, bankAccount.AccountType, bankAccountFromDB.AccountType)
}

// TestCreateBankAccountNoDuplicate tests that the Create method of the BankAccountRepository
// returns an error when trying to create a bank account with the same account number and user.
func TestCreateBankAccountNoDuplicate(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Create a bank account
	ownerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bat.bank, ownerID, "account_number", domain.Checking, "USD")

	// Save the bank account (function to be tested)
	err := bat.repoBA.Create(bankAccount)
	assert.NoError(t, err)

	// Save the bank account again
	err = bat.repoBA.Create(bankAccount)
	assert.Error(t, err)
	assert.Equal(t, model_errors.ErrDbBankAccountAlreadyExists, err)
}

// The bank account must be linked to an existing bank.
func TestCreateBankAccountNoBank(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	bank := domain.NewBank("Test BankN", "https://testbankN.com")
	// Create a bank account
	ownerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bank, ownerID, "account_number", domain.Checking, "USD")

	// Save the bank account (function to be tested)
	err := bat.repoBA.Create(bankAccount)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrForeignKeyViolated, err)
}

func TestCreateBankAccountNoOwner(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Create a bank account
	bankAccount := domain.NewBankAccount(bat.bank, "", "account_number", domain.Checking, "USD")

	// Save the bank account (function to be tested)
	err := bat.repoBA.Create(bankAccount)
	assert.Error(t, err)
	assert.Equal(t, model_errors.ErrDbBankAccountOwnedIDEmpty, err)
}

func TestBankAccountUpdate(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	initialOwnerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bat.bank, initialOwnerID, "account_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccount)

	// Update the bank account
	newOwnerId := uuid.New().String()
	bankAccount.AccountNumber = "new_account_number"
	bankAccount.AccountType = domain.Savings
	bankAccount.Currency = "EUR"
	bankAccount.OwnerID = newOwnerId
	err := bat.repoBA.Update(bankAccount)
	assert.NoError(t, err)
}

func TestBankAccountUpdateNoOwner(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	initialOwnerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bat.bank, initialOwnerID, "account_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccount)

	// Update the bank account
	bankAccount.OwnerID = ""
	err := bat.repoBA.Update(bankAccount)
	assert.Error(t, err)
	assert.Equal(t, model_errors.ErrDbBankAccountOwnedIDEmpty, err)
}

func TestBankAccountUpdateNoBank(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	initialOwnerID := uuid.New().String()
	bank := domain.NewBank("Test BankN", "https://testbankN.com")
	bankAccount := domain.NewBankAccount(bank, initialOwnerID, "account_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccount)

	err := bat.repoBA.Update(bankAccount)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrForeignKeyViolated, err)
}

func TestBankAccountUpdateBadAccountType(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	initialOwnerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bat.bank, initialOwnerID, "account_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccount)

	// Update the bank account
	bankAccount.AccountType = "bad_account_type"
	err := bat.repoBA.Update(bankAccount)
	assert.Error(t, err)
	assert.Equal(t, model_errors.ErrDbBankAccountInvalidType, err)
}

func TestBankAccountDelete(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Create a bank account
	ownerID := uuid.New().String()
	bankAccount := domain.NewBankAccount(bat.bank, ownerID, "account_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccount)

	// Delete the bank account
	err := bat.repoBA.Delete(bankAccount.ID)
	assert.NoError(t, err)

	// Check that the bank account has been deleted
	var bankAccountFromDB *models.BankAccountGormModel
	err = bat.db.First(&bankAccountFromDB, "id = ?", bankAccount.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestBankAccountFindByID(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Create a bank account
	ownerID := uuid.New().String()
	accountID := uuid.New()

	bankAccount := domain.NewBankAccount(bat.bank, ownerID, "account_number", domain.Checking, "USD")
	bankAccount.ID = accountID

	bat.repoBA.Create(bankAccount)

	// Find the bank account by ID
	bankAccountFromDB, err := bat.repoBA.FindByID(accountID)
	assert.NoError(t, err)
	assert.Equal(t, bankAccount.ID, bankAccountFromDB.ID)
	assert.Equal(t, bankAccount.OwnerID, bankAccountFromDB.OwnerID)
	assert.Equal(t, bankAccount.AccountNumber, bankAccountFromDB.AccountNumber)
	assert.Equal(t, bankAccount.AccountType, bankAccountFromDB.AccountType)
}

func TestBankAccountFindByIDNotFound(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Find the bank account by ID
	_, err := bat.repoBA.FindByID(uuid.New())
	assert.Error(t, err)
	assert.Equal(t, model_errors.ErrDbBankAccountNotFound, err)
}

func TestBankAccountFindAll(t *testing.T) {
	bat := NewBankAccountRepositoryTest()

	// Create a bank account
	ownerID := uuid.New().String()
	bankAccountA := domain.NewBankAccount(bat.bank, ownerID, "account_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccountA)

	// Create another bank account
	bankAccountB := domain.NewBankAccount(bat.bank, ownerID, "second_number", domain.Checking, "USD")
	bat.repoBA.Create(bankAccountB)

	// Find all bank accounts
	bankAccounts, err := bat.repoBA.FindAll()
	assert.NoError(t, err)
	assert.Len(t, bankAccounts, 2)

	// Check the first bank account
	assert.Equal(t, bankAccountA.ID, bankAccounts[0].ID)
	assert.Equal(t, bankAccountA.OwnerID, bankAccounts[0].OwnerID)
	assert.Equal(t, bankAccountA.AccountNumber, bankAccounts[0].AccountNumber)
	assert.Equal(t, bankAccountA.AccountType, bankAccounts[0].AccountType)

	// Check the second bank account
	assert.Equal(t, bankAccountB.ID, bankAccounts[1].ID)
	assert.Equal(t, bankAccountB.OwnerID, bankAccounts[1].OwnerID)
	assert.Equal(t, bankAccountB.AccountNumber, bankAccounts[1].AccountNumber)
	assert.Equal(t, bankAccountB.AccountType, bankAccounts[1].AccountType)
}
