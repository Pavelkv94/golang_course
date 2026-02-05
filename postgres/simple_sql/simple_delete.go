package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteTask(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	DELETE FROM tasks WHERE id = $1
	`

	_, err := conn.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}
	return nil
}