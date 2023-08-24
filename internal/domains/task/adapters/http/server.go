package http

import (
	"database/sql"

	"github.com/professionsforall/hexagonal-template/internal/domains/task/adapters/repository"
	"github.com/professionsforall/hexagonal-template/internal/domains/task/core/usecase"
	"github.com/professionsforall/hexagonal-template/pkg/httpserver"

	"github.com/gofiber/fiber/v2"
	"github.com/professionsforall/hexagonal-template/pkg/config"
)

func Init(db *sql.DB) {
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskHandler(taskRepository)
	taskController := NewTaskHttpController(taskUseCase)
	app := fiber.New(fiber.Config{
		AppName: config.AppConfig.App.AppName,
	})
	middlewareApply(app)
	registerRoutes(app, taskController)
	httpserver.HttpServer.Add("task", app)
}
