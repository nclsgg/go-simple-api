package controller

import (
	"FirstAPI/internal/api/customError"
	"FirstAPI/internal/api/dto"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (e *EmailController) UpdateEmail(ctx *fiber.Ctx) error {
	updateDTO := dto.UpdateEmailDTO{}
	updateDTO.ID = ctx.Params("id")
	if err := ctx.BodyParser(&updateDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	if errs := e.myValidator.Validate(updateDTO); len(errs) > 0 && errs[0].Error {
		errMessages := make([]string, len(errs))
		for i, err := range errs {
			errMessages[i] = err.Message
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errMessages, " and "),
		}
	}

	response, err := e.emailUseCase.Update(updateDTO.ID, updateDTO.Email, updateDTO.Status)
	if err != nil {
		if errors.Is(err, customError.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"error":   "Email not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.JSON(response)
}
