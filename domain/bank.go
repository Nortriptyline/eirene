package domain

import "github.com/google/uuid"

type Bank struct {
	ID           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Website      string         `json:"website"`
	BankAccounts []*BankAccount `json:"bank_accounts"`
}

func NewBank(name, website string) *Bank {
	return &Bank{
		ID:           uuid.New(),
		Name:         name,
		Website:      website,
		BankAccounts: []*BankAccount{},
	}
}

func (b *Bank) AddNewBankAccount(ownerID, accountNumber string, accountType AccountType, currency string) *BankAccount {
	account := NewBankAccount(b, ownerID, accountNumber, accountType, currency)
	b.BankAccounts = append(b.BankAccounts, account)
	return account
}

func (b *Bank) AddBankAccount(account *BankAccount) {
	b.BankAccounts = append(b.BankAccounts, account)
}
