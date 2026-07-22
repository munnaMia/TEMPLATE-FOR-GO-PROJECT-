package infra

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/munnaMia/ahlan/internal/config"
)

// return a dsn string for postgres
func getConnectionString(cnf *config.Configuration) string {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&application_name=%s",
		cnf.Database.DB_User,
		cnf.Database.DB_Password,
		cnf.Database.DB_Host,
		cnf.Database.DB_Port,
		cnf.Database.DB_Name,
		cnf.Database.SSL_Mode,
		cnf.Service.Name,
	)

	return dsn
}

// return a new db connection pool
func NewConnection(ctx context.Context, cnf *config.Configuration) (*sql.DB, error) {
	dsn := getConnectionString(cnf)

	pool, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// verify the db connection to check it alive or not
	if err := pool.PingContext(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping the database, %w", err)
	}

	pool.SetMaxIdleConns(6) // max idle connection
	pool.SetMaxOpenConns(6) // max open connection default is unlimited

	return pool, nil
}
