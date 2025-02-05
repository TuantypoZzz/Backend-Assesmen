package service

import (
	"context"
	"errors"
	"farhan/configuration"
	"farhan/entity"
	"farhan/model"
	"farhan/repository"
	"farhan/validation"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type transactionWithdrawServiceImpl struct {
	repository.TransactionWithdrawRepository
	repository.AccountRepository
	logger *logrus.Logger
}

func NewTransactionWithdrawService(transactionWithdrawRepository *repository.TransactionWithdrawRepository, accountRepository *repository.AccountRepository) TransactionWithdrawService {
	logger := configuration.NewLogger()
	return &transactionWithdrawServiceImpl{
		TransactionWithdrawRepository: *transactionWithdrawRepository,
		AccountRepository:             *accountRepository,
		logger:                        logger,
	}
}

func (service *transactionWithdrawServiceImpl) Create(ctx context.Context, transactionModel model.CreateTransaction) (model.Transaction, error) {

	validation.Validate(transactionModel)

	account, err := service.AccountRepository.FindByAccountNumber(ctx, transactionModel.AccountNumber)
	if err != nil {
		service.logger.Error("transaction_withdraw_service_impl_FindByAccountNumber_err", logrus.Fields{"data": account, "error": err.Error()})
		return model.Transaction{}, err
	}

	if transactionModel.Amount > account.Balance {
		return model.Transaction{}, errors.New("insufficient balance")
	}
	finalBalance := account.Balance - transactionModel.Amount

	withdraw := entity.Transaction{
		TransactionID: uuid.New(),
		AccountNumber: transactionModel.AccountNumber,
		Type:          "withdraw",
		Amount:        transactionModel.Amount,
		FinalBalance:  finalBalance,
	}

	service.logger.Info("transaction_withdraw_service_impl_create_entity", logrus.Fields{
		"AccountNumber": withdraw.AccountNumber,
		"FinalBalance":  withdraw.FinalBalance,
	})

	_, err = service.TransactionWithdrawRepository.Create(ctx, withdraw)
	if err != nil {
		service.logger.Error("transaction_withdraw_service_impl_create_err", logrus.Fields{"error": err.Error()})
		return model.Transaction{}, err
	}

	err = service.TransactionWithdrawRepository.UpdateBalance(ctx, transactionModel.AccountNumber, finalBalance)
	if err != nil {
		service.logger.Error("transaction_withdraw_service_impl_create_err", logrus.Fields{"error": err.Error()})
		return model.Transaction{}, err
	}

	service.logger.Info("Withdraw transaction completed", logrus.Fields{
		"AccountNumber": withdraw.AccountNumber,
		"FinalBalance":  withdraw.FinalBalance,
	})

	return model.Transaction{
		Saldo: withdraw.FinalBalance,
	}, nil
}
