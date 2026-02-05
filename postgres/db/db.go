package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)


func CreateConnection(ctx context.Context, databaseUrl string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, databaseUrl)
}