package service

import (
	"context"
	"farhan/model"
)

type AccountService interface {
	Create(ctx context.Context, AccountModel model.CreateAccount) model.RegisterAccount
	IsIdentityCardUsed(ctx context.Context, identityCard string) bool
	IsPhoneUsed(ctx context.Context, phone string) bool
	// FindByAccountNumber(ctx context.Context, accountNumber string) bool
}
