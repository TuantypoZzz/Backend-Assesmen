package repository

import (
	"context"
	"farhan/entity"
)

type TransactionDepositRepository interface {
	Create(ctx context.Context, TrasactionDeposit entity.Transaction) (entity.Transaction, error)
	UpdateBalance(ctx context.Context, accountNumber string, newBalance float64) error
}
