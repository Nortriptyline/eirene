package service

import (
	"github.com/Nortriptyline/Eirene/domain/iservice"
	"github.com/Nortriptyline/Eirene/domain/repository"
	"go.uber.org/zap"
)

type BankServiceImpl struct {
	bankRepo repository.IBankRepository
	logger   *zap.SugaredLogger
}

func NewBankService(
	bankRepo repository.IBankRepository,
	logger *zap.SugaredLogger,
) iservice.IBankService {
	return &BankServiceImpl{
		bankRepo: bankRepo,
		logger:   logger,
	}
}

func (s *BankServiceImpl) CreateBank(name string, website string) error {

	return nil
}
