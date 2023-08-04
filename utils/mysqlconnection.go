package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func getMysqlDsn(user, password, host, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
}

func GetMysqlConnection(ctx context.Context, user, password, host, port, database string, timeout time.Duration) (*sql.DB, error) {
	dsn := getMysqlDsn(user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	err = db.PingContext(ctxWithTimeout)
	if err != nil {
		return nil, err
	}

	return db, nil
}
