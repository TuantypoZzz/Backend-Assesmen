package repository

import (
	"context"
	"farhan/entity"
	"farhan/exception"

	"gorm.io/gorm"
)

type balanceRepositoryImpl struct {
	*gorm.DB
}

func NewBalanceRepositoryImpl(DB *gorm.DB) BalanceRepository {
	return &balanceRepositoryImpl{DB: DB}
}

func (b *balanceRepositoryImpl) FindByAccountNumber(ctx context.Context, accountNumber string) (entity.Account, error) {
	var balance entity.Account
	result := b.DB.WithContext(ctx).Where("account_number = ?", accountNumber).First(&balance)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "No Rekening tidak ditemukan",
		})
	}
	return balance, nil
}
