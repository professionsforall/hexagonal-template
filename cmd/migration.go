package cmd

import (
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/professionsforall/hexagonal-template/pkg/config"
	"github.com/professionsforall/hexagonal-template/pkg/log"
	"github.com/professionsforall/hexagonal-template/utils"
	"github.com/spf13/cobra"
)

var migrationCommand = &cobra.Command{
	Use:   "migrate",
	Short: "migrates schema to mysql database",
	Run: func(cmd *cobra.Command, args []string) {
		mysqlConfig := config.AppConfig.Databases.Mysql
		db, err := utils.GetMysqlConnection(
			cmd.Context(),
			mysqlConfig.UserName,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Database,
			time.Second*10,
		)
		if err != nil {
			log.Logger.Error(err)
			cmd.PrintErr(err)
			return
		}
		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			log.Logger.Error(err)

			cmd.PrintErrln(err)
			return
		}

		migration, err := migrate.NewWithDatabaseInstance("file://migrations/mysql", mysqlConfig.Database, driver)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err = migration.Up()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
		log.Logger.Info("Migrated successfully")
	},
}

func init() {
	rootCommand.AddCommand(migrationCommand)
}
