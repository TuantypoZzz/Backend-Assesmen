package service

import (
	"context"
	"farhan/model"
)

type BalanceService interface {
	FindByAccountNumber(c context.Context, accountNumber string) model.Transaction
}
