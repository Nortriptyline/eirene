package model_errors

import "errors"

var (
	// ErrBankAccountNotFound is returned when a bank account is not found in the database.
	ErrDbBankAccountNotFound           = errors.New("err_db_bank_account_not_found")
	ErrDbBankAccountAlreadyExists      = errors.New("err_db_bank_account_already_exists")
	ErrDbBankAccountBankNotFound       = errors.New("err_db_bank_account_bank_not_found")
	ErrDbBankAccountOwnerNotFound      = errors.New("err_db_bank_account_owner_not_found")
	ErrDbBankAccountInvalidType        = errors.New("err_db_bank_account_invalid_type")
	ErrDbBankAccountCreateError        = errors.New("err_db_bank_account_create_error")
	ErrDbBankAccountOwnedIDEmpty       = errors.New("err_db_bank_account_owner_id_empty")
	ErrDbBankAccountAccountNumberEmpty = errors.New("err_db_bank_account_account_number_empty")
	ErrDbBankAccountCurrencyInvalid    = errors.New("err_db_bank_account_currency_invalid")
	ErrDbBankAccountOwnerIDInvalid     = errors.New("err_db_bank_account_owner_id_invalid")
)
