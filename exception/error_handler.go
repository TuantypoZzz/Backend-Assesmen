package exception

import (
	"encoding/json"
	"farhan/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:   400,
			Remark: "Bad Request",
			Data:   messages,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:   404,
			Remark: "Not Found",
			Data:   err.Error(),
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Code:   401,
			Remark: "Unauthorized",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Code:   500,
		Remark: "General Error",
		Data:   err.Error(),
	})
}
