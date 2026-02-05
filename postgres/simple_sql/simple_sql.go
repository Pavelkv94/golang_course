package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)


func CreateTable( ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL UNIQUE,
		description VARCHAR(1000) NOT NULL,
		is_completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		done_at TIMESTAMP DEFAULT NULL
	)
	`
	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}