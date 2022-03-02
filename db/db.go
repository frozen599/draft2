package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func InitDB(dsn string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
