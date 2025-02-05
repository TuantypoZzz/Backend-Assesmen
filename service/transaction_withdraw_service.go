package service

import (
	"context"
	"farhan/model"
)

type TransactionWithdrawService interface {
	Create(ctx context.Context, AccountModel model.CreateTransaction) (model.Transaction, error)
}
