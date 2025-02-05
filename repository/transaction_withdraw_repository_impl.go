package repository

import (
	"context"
	"farhan/entity"
	"farhan/exception"

	"gorm.io/gorm"
)

type transactionWithdrawRepositoryImpl struct {
	*gorm.DB
}

func NewTransactionWithdrawRepositoryImpl(DB *gorm.DB) TransactionWithdrawRepository {
	return &transactionWithdrawRepositoryImpl{DB: DB}
}

func (repository *transactionWithdrawRepositoryImpl) Create(ctx context.Context, TrasactionWithdraw entity.Transaction) (entity.Transaction, error) {
	err := repository.DB.WithContext(ctx).Create(&TrasactionWithdraw).Error
	exception.PanicLogging(err)
	return TrasactionWithdraw, nil
}

func (repository *transactionWithdrawRepositoryImpl) UpdateBalance(ctx context.Context, accountNumber string, newBalance float64) error {
	err := repository.DB.WithContext(ctx).
		Model(&entity.Account{}).
		Where("account_number = ?", accountNumber).
		Update("balance", newBalance).
		Error
	return err
}
