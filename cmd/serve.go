package cmd

import (
	"context"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/professionsforall/hexagonal-template/internal/domains/task/adapters/http"
	"github.com/professionsforall/hexagonal-template/pkg/config"
	"github.com/professionsforall/hexagonal-template/pkg/httpserver"
	"github.com/professionsforall/hexagonal-template/pkg/log"
	"github.com/professionsforall/hexagonal-template/utils"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "serves application at given port",
	Run: func(cmd *cobra.Command, args []string) {

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
		defer func() {
			log.Logger.Info("closing mysql connection")
			err = conn.Close()
			if err != nil {
				log.Logger.Error(err)
			}
		}()
		app := fiber.New(fiber.Config{AppName: config.AppConfig.App.AppName})
		httpserver.Apply(app, config.AppConfig.App.AppPort, log.Logger)

		http.Init(conn)
		wg := &sync.WaitGroup{}
		utils.Launch(log.Logger, wg, httpserver.HttpServer)
	},
}

func init() {
	rootCommand.AddCommand(serveCommand)
}
