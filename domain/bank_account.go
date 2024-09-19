package domain

import (
	"github.com/google/uuid"
)

type AccountType string

const (
	Checking         AccountType = "checking"
	Savings          AccountType = "savings"
	CreditCard       AccountType = "credit_card"
	Loan             AccountType = "loan"
	Investment       AccountType = "investment"
	Mortgage         AccountType = "mortgage"
	Business         AccountType = "business"
	Joint            AccountType = "joint"
	FixedDeposit     AccountType = "fixed_deposit"
	RecurringDeposit AccountType = "recurring_deposit"
)

type BankAccount struct {
	ID            uuid.UUID      `json:"id"`
	Bank          *Bank          `json:"bank"`
	OwnerID       string         `json:"owner"`
	AccountNumber string         `json:"account_number"`
	AccountType   AccountType    `json:"account_type"`
	Transactions  []*Transaction `json:"transactions"`
	Balance       float64        `json:"balance"`
	Currency      string         `json:"currency"`
}

func NewBankAccount(bank *Bank, ownerID string, accountNumber string, accountType AccountType, currency string) *BankAccount {
	if bank == nil {
		return nil
	}

	return &BankAccount{
		ID:            uuid.New(),
		Bank:          bank,
		OwnerID:       ownerID,
		AccountNumber: accountNumber,
		AccountType:   accountType,
		Transactions:  []*Transaction{},
		Balance:       0.0,
		Currency:      currency,
	}
}

func (ba *BankAccount) AddTransaction(transaction *Transaction) {
	ba.Transactions = append(ba.Transactions, transaction)
}

func (ba *BankAccount) CreateTransaction(amount float64, transactionType TransactionType, description string, date string) *Transaction {
	transaction := NewTransaction(ba, amount, transactionType, description, date)
	ba.AddTransaction(transaction)
	return transaction
}

func (t AccountType) IsValid() bool {
	switch t {
	case Checking, Savings, CreditCard, Loan, Investment, Mortgage, Business, Joint, FixedDeposit, RecurringDeposit:
		return true
	}
	return false
}
