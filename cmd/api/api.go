package api

import (
	"FirstAPI/internal/api/controller"
	"FirstAPI/internal/api/middleware"
	"FirstAPI/internal/api/repository"
	"FirstAPI/internal/api/useCase"
	"FirstAPI/internal/config"
	"FirstAPI/internal/infra/db"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	dbConnector *db.MongoDBService
}

func NewApi(dbConnector *db.MongoDBService) *Api {
	return &Api{
		dbConnector: dbConnector,
	}
}

func (a *Api) Start() {
	server := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(middleware.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	emailRepository := repository.NewEmailRepository(a.dbConnector)
	emailUseCase := useCase.NewEmailUseCase(emailRepository)
	emailController := controller.NewEmailController(emailUseCase)

	server.Get("/email", emailController.GetEmails)
	server.Get("/email/:id", emailController.GetByID)
	server.Post("/email", emailController.CreateEmail)
	server.Patch("/email/:id", emailController.UpdateEmail)
	server.Delete("/email/:id", emailController.DeleteEmail)

	err := server.Listen(config.PORT)
	if err != nil {
		panic(err)
	}
}
