package controller

import (
	"farhan/configuration"
	"farhan/model"
	"farhan/repository"
	"farhan/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type balanceController struct {
	service.BalanceService
	repository.AccountRepository
	logger *logrus.Logger
}

func NewBalanceController(balanceService *service.BalanceService, accountRepository *repository.AccountRepository) *balanceController {
	logger := configuration.NewLogger()
	return &balanceController{
		BalanceService:    *balanceService,
		AccountRepository: *accountRepository,
		logger:            logger,
	}
}

func (controller balanceController) Route(api fiber.Router) {
	Balance := api.Group("/saldo")
	Balance.Get("/:account_number", controller.FindByAccountNumber)
}

func (controller balanceController) FindByAccountNumber(ctx *fiber.Ctx) error {
	accountNumber := ctx.Params("account_number")

	controller.logger.Info("balance_controller_FindByAccountNumber_Params", logrus.Fields{
		"data": accountNumber,
	})

	_, errs := controller.AccountRepository.FindByAccountNumber(ctx.Context(), accountNumber)
	if errs != nil {
		controller.logger.Warning("Account number does not exist", logrus.Fields{
			"data": accountNumber,
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "Account number not exist",
		})
	}

	result := controller.BalanceService.FindByAccountNumber(ctx.Context(), accountNumber)

	controller.logger.Info("balance_controller_result", logrus.Fields{
		"data": result,
	})

	return ctx.JSON(model.GeneralResponse{
		Code:   200,
		Remark: "Success",
		Data:   result,
	})
}
