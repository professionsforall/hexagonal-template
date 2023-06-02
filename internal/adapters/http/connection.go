package http

import (
	"database/sql"

	"github.com/professionsforall/hexagonal-template/pkg/config"
	"github.com/professionsforall/hexagonal-template/utils"
)

var databaseConnection = func() (*sql.DB, error) {
	connectionConfig := config.AppConfig.Database.Postgres
	return utils.PostgresConnection(
		connectionConfig.Host,
		connectionConfig.UserName,
		connectionConfig.Password,
		connectionConfig.Port,
		connectionConfig.Database,
		connectionConfig.Sslmode,
	)
}
