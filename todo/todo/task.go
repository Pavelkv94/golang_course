package todo

import "time"

type Task struct {
	Title string
	Description string
	IsCompleted bool

	CreatedAt time.Time
	DoneAt *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title: title,
		Description: description,
		IsCompleted: false,
		CreatedAt: time.Now(),
		DoneAt: nil,
	}
}

func (t *Task) CompleteTask(status bool) {
	doneTime := time.Now()
	t.IsCompleted = status
	t.DoneAt = &doneTime
}