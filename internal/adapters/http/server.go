package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/professionsforall/hexagonal-template/internal/adapters/repository"
	"github.com/professionsforall/hexagonal-template/internal/core/usecase"
	"github.com/professionsforall/hexagonal-template/pkg/config"
	"github.com/professionsforall/hexagonal-template/pkg/log"
	"github.com/professionsforall/hexagonal-template/utils"
)

var BootTaskController TaskController

func Init(ctx context.Context) {
	mysqlConfig := config.AppConfig.Databases.Mysql
	conn, err := utils.GetMysqlConnection(
		context.Background(),
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database,
		mysqlConfig.Timeout,
	)
	if err != nil {
		log.Logger.Panic(err)
	}
	defer conn.Close()

	taskRepository := repository.NewTaskRepository(conn)
	taskUseCase := usecase.NewTaskHandler(taskRepository)
	taskController := NewTaskHttpController(taskUseCase)
	app := fiber.New(fiber.Config{
		AppName:      config.AppConfig.App.AppName,
		ErrorHandler: errorHandler,
	})
	ctxWithTimeout, cancel := context.WithTimeout(ctx, config.AppConfig.ShutdownTime)
	defer cancel()
	defer func() {
		log.Logger.Info("shutting down")
		err = app.ShutdownWithContext(ctxWithTimeout)
		if err != nil {
			log.Logger.Info(err.Error())
		}
	}()
	middlewareApply(app)
	registerRoutes(app, taskController)

	go func() {
		err = app.Listen(":" + config.AppConfig.App.AppPort)
		if err != nil {
			log.Logger.Panic(err)
		}
	}()
	utils.Notify()
	go func() {
		<-ctxWithTimeout.Done()
		if ctx.Err() != nil {
			log.Logger.Error(ctxWithTimeout.Err())
		}
	}()
}
