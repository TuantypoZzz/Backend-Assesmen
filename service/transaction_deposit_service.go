package service

import (
	"context"
	"farhan/model"
)

type TransactionDepositService interface {
	Create(ctx context.Context, AccountModel model.CreateTransaction) (model.Transaction, error)
}
