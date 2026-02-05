package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TaskDto struct {
	Title string
	Description string
}

func InsertTask(ctx context.Context, conn *pgx.Conn, taskDto TaskDto) error {
	sqlQuery := `
	INSERT INTO tasks (title, description) VALUES ($1, $2)
	`

	_, err := conn.Exec(ctx, sqlQuery, taskDto.Title, taskDto.Description)
	if err != nil {
		return err
	}
	return nil
}