package simplesql

import "time"

type TaskModel struct {
	Id int
	Title string
	Description string
	IsCompleted bool
	CreatedAt time.Time
	DoneAt *time.Time
}

