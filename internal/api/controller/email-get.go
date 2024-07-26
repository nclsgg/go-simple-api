package controller

import "github.com/gofiber/fiber/v2"

func (e *EmailController) GetEmails(ctx *fiber.Ctx) error {
	emails, _ := e.emailUseCase.GetEmails()

	return ctx.JSON(emails)
}
