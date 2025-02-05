package repository

import (
	"context"
	"farhan/entity"
	"farhan/exception"

	"gorm.io/gorm"
)

type transactionDepositRepositoryImpl struct {
	*gorm.DB
}

func NewTransactionDepositRepositoryImpl(DB *gorm.DB) TransactionDepositRepository {
	return &transactionDepositRepositoryImpl{DB: DB}
}

func (repository *transactionDepositRepositoryImpl) Create(ctx context.Context, TrasactionDeposit entity.Transaction) (entity.Transaction, error) {
	err := repository.DB.WithContext(ctx).Create(&TrasactionDeposit).Error
	exception.PanicLogging(err)
	return TrasactionDeposit, nil
}

func (repository *transactionDepositRepositoryImpl) UpdateBalance(ctx context.Context, accountNumber string, newBalance float64) error {
	err := repository.DB.WithContext(ctx).
		Model(&entity.Account{}).
		Where("account_number = ?", accountNumber).
		Update("balance", newBalance).
		Error
	return err
}
