package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UpdateTaskDto struct {
	Id int
	Title string
	Description string
}

func UpdateTask(ctx context.Context, conn *pgx.Conn, taskDto UpdateTaskDto) error {
	sqlQuery := `
	UPDATE tasks SET title = $1, description = $2 WHERE id = $3
	`

	_, err := conn.Exec(ctx, sqlQuery, taskDto.Title, taskDto.Description, taskDto.Id)
	if err != nil {
		return err
	}
	return nil
}