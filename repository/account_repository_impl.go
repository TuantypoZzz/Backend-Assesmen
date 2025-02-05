package repository

import (
	"context"
	"farhan/entity"
	"farhan/exception"

	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	*gorm.DB
}

func NewAccountRepositoryImpl(DB *gorm.DB) AccountRepository {
	return &accountRepositoryImpl{DB: DB}
}

func (repository *accountRepositoryImpl) Create(ctx context.Context, account entity.Account) (entity.Account, error) {
	err := repository.DB.WithContext(ctx).Create(&account).Error
	exception.PanicLogging(err)
	return account, nil
}

func (repository *accountRepositoryImpl) IsIdentityCardUsed(ctx context.Context, identityCard string) bool {
	var count int64
	repository.DB.WithContext(ctx).
		Model(&entity.Account{}).
		Where("identity_card = ?", identityCard).
		Count(&count)
	return count > 0
}

func (repository *accountRepositoryImpl) IsPhoneUsed(ctx context.Context, phone string) bool {
	var count int64
	repository.DB.WithContext(ctx).
		Model(&entity.Account{}).
		Where("phone = ?", phone).
		Count(&count)
	return count > 0
}

func (repository *accountRepositoryImpl) FindByAccountNumber(ctx context.Context, accountNumber string) (entity.Account, error) {
	var account entity.Account
	result := repository.DB.WithContext(ctx).Where("account_number = ?", accountNumber).First(&account)
	if result.Error != nil {
		//log error
		return account, result.Error
	}
	return account, nil
}
