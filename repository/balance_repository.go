package repository

import (
	"context"
	"farhan/entity"
)

type BalanceRepository interface {
	FindByAccountNumber(ctx context.Context, accountNumber string) (entity.Account, error)
}
