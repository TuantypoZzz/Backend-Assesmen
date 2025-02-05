package service

import (
	"context"
	"farhan/configuration"
	"farhan/entity"
	"farhan/exception"
	"farhan/model"
	"farhan/repository"
	"farhan/validation"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type accountServiceImpl struct {
	repository.AccountRepository
	logger *logrus.Logger
}

func NewAccountServiceImpl(accountRepository *repository.AccountRepository) AccountService {
	logger := configuration.NewLogger() // Menggunakan logger terstruktur
	return &accountServiceImpl{
		AccountRepository: *accountRepository,
		logger:            logger,
	}
}

func (service *accountServiceImpl) Create(ctx context.Context, accountModel model.CreateAccount) model.RegisterAccount {

	validation.Validate(accountModel)

	account := entity.Account{
		AccountID:     uuid.New(),
		AccountNumber: generateAccountNumber(),
		AccountName:   accountModel.AccountName,
		IdentityCard:  accountModel.IdentityCard,
		Phone:         accountModel.Phone,
		CreatedAt:     time.Now(),
	}

	service.logger.Info("account_service_impl_create_entity", logrus.Fields{
		"AccountID":     account.AccountID,
		"AccountNumber": account.AccountNumber,
	})

	account, err := service.AccountRepository.Create(ctx, account)
	if err != nil {
		service.logger.Error("account_service_impl_create_err", logrus.Fields{
			"error": err.Error(),
		})
		exception.PanicLogging(err)
	}

	return model.RegisterAccount{
		AccountNumber: account.AccountNumber,
	}
}

func (service *accountServiceImpl) IsIdentityCardUsed(ctx context.Context, identityCard string) bool {
	return service.AccountRepository.IsIdentityCardUsed(ctx, identityCard)
}

func (service *accountServiceImpl) IsPhoneUsed(ctx context.Context, phone string) bool {
	return service.AccountRepository.IsPhoneUsed(ctx, phone)
}

func generateAccountNumber() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
