package main

import (
	"farhan/configuration"
	"farhan/controller"
	"farhan/exception"
	"farhan/repository"
	"farhan/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// configuration
	database := configuration.NewDatabase()

	// repository
	accountRepository := repository.NewAccountRepositoryImpl(database)
	transactionDepositRepository := repository.NewTransactionDepositRepositoryImpl(database)
	transactionWithdrawRepository := repository.NewTransactionWithdrawRepositoryImpl(database)
	balanceRepository := repository.NewBalanceRepositoryImpl(database)

	// services
	accountService := service.NewAccountServiceImpl(&accountRepository)
	transactionDepositService := service.NewTransactionDepositService(&transactionDepositRepository, &accountRepository)
	transactionWithdrawService := service.NewTransactionWithdrawService(&transactionWithdrawRepository, &accountRepository)
	balanceService := service.NewBalanceService(&balanceRepository)

	// controller
	accountController := controller.NewAccountController(&accountService)
	transactionDepositController := controller.NewTransactionDepositController(&transactionDepositService, &accountRepository)
	transactionWithdrawController := controller.NewTransactionWithdrawController(&transactionWithdrawService, &accountRepository)
	balanceController := controller.NewBalanceController(&balanceService, &accountRepository)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(configuration.NewLoggerConfig()))

	//routing
	api := app.Group("/api/v1")
	accountController.Route(api)
	transactionDepositController.Route(api)
	transactionWithdrawController.Route(api)
	balanceController.Route(api)

	err := app.Listen(os.Getenv("SERVER_PORT"))
	exception.PanicLogging(err)
}
