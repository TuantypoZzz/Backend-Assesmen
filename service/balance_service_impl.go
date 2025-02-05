package service

import (
	"context"
	"farhan/configuration"
	"farhan/model"
	"farhan/repository"

	"github.com/sirupsen/logrus"
)

type balanceServiceImpl struct {
	repository.BalanceRepository
	logger *logrus.Logger
}

func NewBalanceService(balanceRepository *repository.BalanceRepository) BalanceService {
	logger := configuration.NewLogger()
	return &balanceServiceImpl{BalanceRepository: *balanceRepository, logger: logger}
}

func (b *balanceServiceImpl) FindByAccountNumber(c context.Context, accountNumber string) model.Transaction {
	balance, err := b.BalanceRepository.FindByAccountNumber(c, accountNumber)
	if err != nil {
		b.logger.Error("balance_service_impl_FindByAccountNumber_err", logrus.Fields{
			"error": err.Error(),
		})
		return model.Transaction{}
	}

	b.logger.Info("balance_service_impl_FindByAccountNumber_result", logrus.Fields{
		"data": accountNumber,
	})

	return model.Transaction{
		Saldo: balance.Balance,
	}
}
