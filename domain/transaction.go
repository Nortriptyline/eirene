package domain

import "github.com/google/uuid"

type TransactionType string

const (
	Deposit       TransactionType = "deposit"
	Withdrawal    TransactionType = "withdrawal"
	Transfer      TransactionType = "transfer"
	Payment       TransactionType = "payment"
	Fee           TransactionType = "fee"
	Interest      TransactionType = "interest"
	Refund        TransactionType = "refund"
	Purchase      TransactionType = "purchase"
	LoanRepayment TransactionType = "loan_repayment"
	Dividend      TransactionType = "dividend"
)

type Transaction struct {
	ID              uuid.UUID       `json:"id"`
	BankAccountID   uuid.UUID       `json:"bank_account_id"`
	BankAccount     *BankAccount    `json:"bank_account"`
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"transaction_type"`
	Description     string          `json:"description"`
	TransactionDate string          `json:"transaction_date"`
}

func NewTransaction(bankAccount *BankAccount, amount float64, transactionType TransactionType, description, transactionDate string) *Transaction {
	return &Transaction{
		ID:              uuid.New(),
		BankAccount:     bankAccount,
		Amount:          amount,
		TransactionType: transactionType,
		Description:     description,
		TransactionDate: transactionDate,
	}
}
