package controller

import (
	"farhan/configuration"
	"farhan/model"
	"farhan/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AccountController struct {
	service.AccountService
	logger *logrus.Logger
}

func NewAccountController(accountService *service.AccountService) *AccountController {
	logger := configuration.NewLogger() // Menggunakan logger
	return &AccountController{
		AccountService: *accountService,
		logger:         logger,
	}
}

func (controller AccountController) Route(api fiber.Router) {
	Account := api.Group("/daftar")
	Account.Post("/", controller.Create)
}

func (controller AccountController) Create(c *fiber.Ctx) error {
	var request model.CreateAccount
	err := c.BodyParser(&request)
	if err != nil {
		controller.logger.Error("account_controller_create_bodyParser", logrus.Fields{
			"error": err.Error(),
		})
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "format salah",
		})
	}

	controller.logger.Info("account_controller_create_request", logrus.Fields{
		"data": request,
	})

	// Validasi apakah identity_card sudah digunakan
	if controller.AccountService.IsIdentityCardUsed(c.Context(), request.IdentityCard) {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "NIK sudah digunakan",
		})
	}

	// Validasi apakah phone sudah digunakan
	if controller.AccountService.IsPhoneUsed(c.Context(), request.Phone) {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "No HP sudah digunakan",
		})
	}

	response := controller.AccountService.Create(c.Context(), request)
	if response.AccountNumber == "" {
		controller.logger.Error("Failed to create account")
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:   500,
			Remark: "Gagal membuat account",
		})
	}

	controller.logger.Info("Account successfully created", logrus.Fields{
		"AccountNumber": response.AccountNumber,
	})

	return c.JSON(model.GeneralResponse{
		Code:   200,
		Remark: "Data has been created",
		Data:   response,
	})
}
