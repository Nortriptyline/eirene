package bankaccount

import "errors"

var (
	ErrBankAccountBankIsNil          = errors.New("bank_is_nil")
	ErrBankAccountOwnedIDEmpty       = errors.New("owner_id_empty")
	ErrBankAccountAccountNumberEmpty = errors.New("account_number_empty")
	ErrBankAccountCurrencyInvalid    = errors.New("currency_invalid")
	ErrBankAccountOwnerIDInvalid     = errors.New("owner_id_invalid")
)
