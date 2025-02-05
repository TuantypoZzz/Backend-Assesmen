package repository

import (
	"context"
	"farhan/entity"
)

type AccountRepository interface {
	Create(ctx context.Context, account entity.Account) (entity.Account, error)
	IsIdentityCardUsed(ctx context.Context, identityCard string) bool
	IsPhoneUsed(ctx context.Context, phone string) bool
	FindByAccountNumber(ctx context.Context, accountNumber string) (entity.Account, error)
}
