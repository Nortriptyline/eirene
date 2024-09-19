package application

import (
	"github.com/Nortriptyline/Eirene/application/service"
	"github.com/Nortriptyline/Eirene/domain/iservice"
	"go.uber.org/zap"
)

type Services struct {
	BankService iservice.IBankService
	// BankAccountService BankAccountService
	// TransactionService TransactionService
}

func InitializeServices(repos *Repositories, logger *zap.SugaredLogger) *Services {
	return &Services{
		BankService: service.NewBankService(repos.BankRepository, logger),
		// BankAccountService: NewBankAccountService(repos.BankAccountRepository),
		// TransactionService: NewTransactionService(repos.TransactionRepository),
	}
}
