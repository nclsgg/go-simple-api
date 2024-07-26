package controller

import (
	"FirstAPI/internal/api/customError"
	"FirstAPI/internal/api/dto"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (e *EmailController) DeleteEmail(ctx *fiber.Ctx) error {
	deleteDTO := dto.DeleteEmailDTO{}
	deleteDTO.ID = ctx.Params("id")

	if err := e.emailUseCase.Delete(deleteDTO.ID); err != nil {
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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Email deleted successfully",
	})
}
