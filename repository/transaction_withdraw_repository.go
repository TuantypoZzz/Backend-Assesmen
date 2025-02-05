package repository

import (
	"context"
	"farhan/entity"
)

type TransactionWithdrawRepository interface {
	Create(ctx context.Context, TrasactionWithdraw entity.Transaction) (entity.Transaction, error)
	UpdateBalance(ctx context.Context, accountNumber string, newBalance float64) error
}
