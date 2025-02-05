package service

import (
	"context"
	"farhan/configuration"
	"farhan/entity"
	"farhan/model"
	"farhan/repository"
	"farhan/validation"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type transactionDepositServiceImpl struct {
	repository.TransactionDepositRepository
	repository.AccountRepository
	logger *logrus.Logger
}

func NewTransactionDepositService(transactionDepositRepository *repository.TransactionDepositRepository, accountRepository *repository.AccountRepository) TransactionDepositService {
	logger := configuration.NewLogger()
	return &transactionDepositServiceImpl{
		TransactionDepositRepository: *transactionDepositRepository,
		AccountRepository:            *accountRepository,
		logger:                       logger,
	}
}

func (service *transactionDepositServiceImpl) Create(ctx context.Context, transactionModel model.CreateTransaction) (model.Transaction, error) {

	validation.Validate(transactionModel)

	account, err := service.AccountRepository.FindByAccountNumber(ctx, transactionModel.AccountNumber)
	if err != nil {
		service.logger.Error("transaction_deposit_service_impl_FindByAccountNumber_err", logrus.Fields{"data": account, "error": err.Error()})
		return model.Transaction{}, err
	}

	finalBalance := account.Balance + transactionModel.Amount

	deposit := entity.Transaction{
		TransactionID: uuid.New(),
		AccountNumber: transactionModel.AccountNumber,
		Type:          "deposit",
		Amount:        transactionModel.Amount,
		FinalBalance:  finalBalance,
	}

	service.logger.Info("transaction_deposit_service_impl_create_entity", logrus.Fields{
		"data": deposit,
	})

	_, err = service.TransactionDepositRepository.Create(ctx, deposit)
	if err != nil {
		service.logger.Error("transaction_deposit_service_impl_create_err", logrus.Fields{"error": err.Error()})
		return model.Transaction{}, err
	}

	err = service.TransactionDepositRepository.UpdateBalance(ctx, transactionModel.AccountNumber, finalBalance)
	if err != nil {
		service.logger.Error("transaction_deposit_service_impl_UpdateBalace_err", logrus.Fields{"error": err.Error()})
		return model.Transaction{}, err
	}

	return model.Transaction{
		Saldo: deposit.FinalBalance,
	}, nil
}
