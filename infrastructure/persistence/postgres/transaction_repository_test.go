package postgres

import (
	"testing"

	"github.com/Nortriptyline/Eirene/domain"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models"
	"github.com/Nortriptyline/Eirene/infrastructure/persistence/models/model_errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TransactionRepositoryTest struct {
	db     *gorm.DB
	bRepo  repository.IBankRepository
	baRepo repository.IBankAccountRepository
	tRepo  repository.ITransactionRepository
}

func NewTransactionRepositoryTest() *TransactionRepositoryTest {
	db := SetupTestDB()
	return &TransactionRepositoryTest{
		db:     db,
		bRepo:  NewBankRepository(db),
		baRepo: NewBankAccountRepository(db),
		tRepo:  NewTransactionRepository(db),
	}
}

func TestTransactionRepository_Create(t *testing.T) {
	trt := NewTransactionRepositoryTest()

	b := domain.NewBank("Test Bank", "https://testbank.com")
	ba := domain.NewBankAccount(b, "owner123", "123456789", domain.Checking, "USD")
	transaction := domain.NewTransaction(ba, 100.0, domain.Deposit, "Test transaction", "2023-10-01")

	err := trt.bRepo.Create(b)
	assert.NoError(t, err)
	err = trt.baRepo.Create(ba)
	assert.NoError(t, err)

	err = trt.tRepo.Create(transaction)
	assert.NoError(t, err)

	var transactionFromDB models.TransactionGormModel
	err = trt.db.First(&transactionFromDB, "id = ?", transaction.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, transaction.ID, transactionFromDB.ID)
	assert.Equal(t, transaction.Amount, transactionFromDB.Amount)
	assert.Equal(t, transaction.TransactionType, transactionFromDB.TransactionType)
	assert.Equal(t, transaction.Description, transactionFromDB.Description)
	assert.Equal(t, transaction.TransactionDate, transactionFromDB.TransactionDate)
}

func TestTransactionRepository_Update(t *testing.T) {
	trt := NewTransactionRepositoryTest()

	b := domain.NewBank("Test Bank", "https://testbank.com")
	ba := domain.NewBankAccount(b, "owner123", "123456789", domain.Checking, "USD")
	transaction := domain.NewTransaction(ba, 100.0, domain.Deposit, "Test transaction", "2023-10-01")

	trt.bRepo.Create(b)
	trt.baRepo.Create(ba)

	err := trt.tRepo.Create(transaction)
	assert.NoError(t, err)

	transaction.Amount = 200.0
	transaction.Description = "Updated transaction"

	err = trt.tRepo.Update(transaction)
	assert.NoError(t, err)

	var transactionFromDB models.TransactionGormModel
	err = trt.db.First(&transactionFromDB, "id = ?", transaction.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, transaction.Amount, transactionFromDB.Amount)
	assert.Equal(t, transaction.Description, transactionFromDB.Description)
}

func TestTransactionRepository_Delete(t *testing.T) {
	trt := NewTransactionRepositoryTest()
	b := domain.NewBank("Test Bank", "https://testbank.com")
	ba := domain.NewBankAccount(b, "owner123", "123456789", domain.Checking, "USD")
	transaction := domain.NewTransaction(ba, 100.0, domain.Deposit, "Test transaction", "2023-10-01")

	trt.bRepo.Create(b)
	trt.baRepo.Create(ba)

	err := trt.tRepo.Create(transaction)
	assert.NoError(t, err)

	err = trt.tRepo.Delete(transaction.ID)
	assert.NoError(t, err)

	var transactionFromDB models.TransactionGormModel
	err = trt.db.First(&transactionFromDB, "id = ?", transaction.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestTransactionRepository_FindByID(t *testing.T) {
	trt := NewTransactionRepositoryTest()
	b := domain.NewBank("Test Bank", "https://testbank.com")
	ba := domain.NewBankAccount(b, "owner123", "123456789", domain.Checking, "USD")
	transaction := domain.NewTransaction(ba, 100.0, domain.Deposit, "Test transaction", "2023-10-01")

	trt.bRepo.Create(b)
	trt.baRepo.Create(ba)

	err := trt.tRepo.Create(transaction)
	assert.NoError(t, err)

	foundTransaction, err := trt.tRepo.FindByID(transaction.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundTransaction)
	assert.Equal(t, transaction.ID, foundTransaction.ID)
	assert.Equal(t, transaction.Amount, foundTransaction.Amount)
	assert.Equal(t, transaction.TransactionType, foundTransaction.TransactionType)
	assert.Equal(t, transaction.Description, foundTransaction.Description)
	assert.Equal(t, transaction.TransactionDate, foundTransaction.TransactionDate)
}

func TestTransactionRepository_FindByID_NotFound(t *testing.T) {
	trt := NewTransactionRepositoryTest()
	id := uuid.New()

	transaction, err := trt.tRepo.FindByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, model_errors.ErrDbTransactionNotFound, err)
}

func TestTransactionRepository_FindAll(t *testing.T) {
	trt := NewTransactionRepositoryTest()
	b := domain.NewBank("Test Bank", "https://testbank.com")
	ba := domain.NewBankAccount(b, "owner123", "123456789", domain.Checking, "USD")
	transaction1 := domain.NewTransaction(ba, 100.0, domain.Payment, "Test transaction 1", "2023-10-01")
	transaction2 := domain.NewTransaction(ba, 200.0, domain.Deposit, "Test transaction 2", "2023-10-02")

	trt.bRepo.Create(b)
	trt.baRepo.Create(ba)

	err := trt.tRepo.Create(transaction1)
	assert.NoError(t, err)
	err = trt.tRepo.Create(transaction2)
	assert.NoError(t, err)

	transactions, err := trt.tRepo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, transactions, 2)

	assert.Equal(t, transaction1.ID, transactions[0].ID)
	assert.Equal(t, transaction1.Amount, transactions[0].Amount)
	assert.Equal(t, transaction1.TransactionType, transactions[0].TransactionType)
	assert.Equal(t, transaction1.Description, transactions[0].Description)
	assert.Equal(t, transaction1.TransactionDate, transactions[0].TransactionDate)

	assert.Equal(t, transaction2.ID, transactions[1].ID)
	assert.Equal(t, transaction2.Amount, transactions[1].Amount)
	assert.Equal(t, transaction2.TransactionType, transactions[1].TransactionType)
	assert.Equal(t, transaction2.Description, transactions[1].Description)
	assert.Equal(t, transaction2.TransactionDate, transactions[1].TransactionDate)
}

// FindAll must return an empty slice if there are no transactions
func TestTransactionRepository_FindAll_Empty(t *testing.T) {
	trt := NewTransactionRepositoryTest()

	transactions, err := trt.tRepo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, transactions, 0)
}
