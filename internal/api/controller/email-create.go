package controller

import (
	"FirstAPI/internal/api/dto"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (e *EmailController) CreateEmail(ctx *fiber.Ctx) error {
	createDTO := &dto.CreateEmailDTO{}
	if err := ctx.BodyParser(&createDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	if errs := e.myValidator.Validate(createDTO); len(errs) > 0 && errs[0].Error {
		errMessages := make([]string, len(errs))
		for i, err := range errs {
			errMessages[i] = err.Message
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errMessages, " and "),
		}
	}

	result, err := e.emailUseCase.Create(createDTO.Email, createDTO.Status)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.JSON(result)
}
