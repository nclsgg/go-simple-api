package controller

import (
	"FirstAPI/internal/api/customError"
	"FirstAPI/internal/api/dto"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (e *EmailController) GetByID(ctx *fiber.Ctx) error {
	findDTO := &dto.FindByIDDTO{}
	findDTO.ID = ctx.Params("id")

	response, err := e.emailUseCase.GetByID(findDTO.ID)
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
