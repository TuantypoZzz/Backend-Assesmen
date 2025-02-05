package controller

import (
	"farhan/configuration"
	"farhan/model"
	"farhan/repository"
	"farhan/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionWithdrawController struct {
	service.TransactionWithdrawService
	repository.AccountRepository
	logger *logrus.Logger
}

func NewTransactionWithdrawController(transactionWithdrawService *service.TransactionWithdrawService, accountRepository *repository.AccountRepository) *TransactionWithdrawController {
	logger := configuration.NewLogger()
	return &TransactionWithdrawController{
		TransactionWithdrawService: *transactionWithdrawService,
		AccountRepository:          *accountRepository,
		logger:                     logger,
	}
}

func (controller TransactionWithdrawController) Route(api fiber.Router) {
	Transaction := api.Group("/tarik")
	Transaction.Post("/", controller.Create)
}

func (controller TransactionWithdrawController) Create(c *fiber.Ctx) error {
	var request model.CreateTransaction
	err := c.BodyParser(&request)
	if err != nil {
		controller.logger.Error("transaction_withdraw_controller_create_bodyParser", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "format salah",
		})
	}

	controller.logger.Info("transaction_withdraw_controller_create_request", logrus.Fields{
		"data": request,
	})

	account, errs := controller.AccountRepository.FindByAccountNumber(c.Context(), request.AccountNumber)
	if errs != nil {
		controller.logger.Warning("transaction_withdraw_controller_FindByAccountNumber_err", logrus.Fields{"data": request})
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "nomor rekening tidak terdaftar",
		})
	}

	if request.Amount > account.Balance {
		controller.logger.Warning("transaction_withdraw_controller_create_err_transactionModel.Amount>account.Balance", logrus.Fields{
			"data": account.Balance,
		})
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "Saldo tidak cukup",
		})
	}

	response, err := controller.TransactionWithdrawService.Create(c.Context(), request)
	if err != nil {
		controller.logger.Error("transaction_withdraw_controller_create_err", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:   500,
			Remark: "Gagal membuat data",
			Data:   nil,
		})
	}

	return c.JSON(model.GeneralResponse{
		Code:   200,
		Remark: "Data telah dibuat",
		Data:   response,
	})
}
