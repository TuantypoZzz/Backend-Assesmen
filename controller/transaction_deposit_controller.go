package controller

import (
	"farhan/configuration"
	"farhan/model"
	"farhan/repository"
	"farhan/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionDepositController struct {
	service.TransactionDepositService
	repository.AccountRepository
	logger *logrus.Logger
}

func NewTransactionDepositController(transactionDepositService *service.TransactionDepositService, accountRepository *repository.AccountRepository) *TransactionDepositController {
	logger := configuration.NewLogger()
	return &TransactionDepositController{
		TransactionDepositService: *transactionDepositService,
		AccountRepository:         *accountRepository,
		logger:                    logger,
	}
}

func (controller TransactionDepositController) Route(api fiber.Router) {
	Transaction := api.Group("/tabung")
	Transaction.Post("/", controller.Create)
}

func (controller TransactionDepositController) Create(c *fiber.Ctx) error {
	var request model.CreateTransaction
	err := c.BodyParser(&request)
	if err != nil {
		controller.logger.Error("transaction_deposit_controller_create_bodyParser", logrus.Fields{"error": err.Error()})
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "format salah",
		})
	}

	controller.logger.Info("transaction_deposit_controller_create_reqeust", logrus.Fields{
		"data": request,
	})

	_, errs := controller.AccountRepository.FindByAccountNumber(c.Context(), request.AccountNumber)
	if errs != nil {
		controller.logger.Warning("transaction_deposit_controller_create_FindByAccountNumber", logrus.Fields{"data": request})
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "nomor rekening tidak terdaftar",
		})
	}

	response, err := controller.TransactionDepositService.Create(c.Context(), request)
	if err != nil {
		controller.logger.Error("transaction_deposit_controller_create_err", logrus.Fields{"error": err.Error()})
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
