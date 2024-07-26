package controller

import (
	"FirstAPI/internal/api/middleware"
	"FirstAPI/internal/api/useCase"
)

type EmailController struct {
	emailUseCase *useCase.EmailUseCase
	myValidator  middleware.XValidator
}

func NewEmailController(useCase *useCase.EmailUseCase) *EmailController {
	return &EmailController{
		emailUseCase: useCase,
		myValidator: middleware.XValidator{
			Validator: middleware.Validate,
		},
	}
}
