package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectTasks(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT * FROM tasks
	ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery) // получаем сырые данные из базы
	if err != nil {
		return nil, err
	}

	defer rows.Close() // закрываем rows чтобы не было утечки ресурсов

	tasks := make([]TaskModel, 0)

	for rows.Next() { // проходим по всем строкам и получаем данные
		var task TaskModel

		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.DoneAt) // сканируем данные из строки
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}