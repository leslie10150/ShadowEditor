// +build ignore

package helper

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// NewPostgreSQL create a new postgreSQL connection
func NewPostgreSQL(host string, port uint16, user, password, database string) (*PostgreSQL, error) {
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil, err
	}
	connConfig.Host = host
	connConfig.Port = port
	connConfig.User = user
	connConfig.Password = password
	connConfig.Database = database

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.TODO())
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{conn}, nil
}

// PostgreSQL postgreSQL connection
type PostgreSQL struct {
	Connection *pgx.Conn
}

// Query query records with sql string
func (p PostgreSQL) Query(sql string, args ...interface{}) (pgx.Rows, error) {
	if p.Connection == nil {
		return nil, fmt.Errorf("connection is not created")
	}

	return p.Connection.Query(context.Background(), sql, args...)
}

// Exec execute a sql
func (p PostgreSQL) Exec(sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if p.Connection == nil {
		return nil, fmt.Errorf("connection is not created")
	}

	return p.Connection.Exec(context.TODO(), sql, args...)
}

// Close close connection
func (p PostgreSQL) Close() error {
	if p.Connection == nil {
		return fmt.Errorf("connection is not created")
	}

	return p.Connection.Close(context.TODO())
}
