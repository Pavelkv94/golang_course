package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task, 0),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock() // блокировка чтобы в случае конкурентного доступа к map не было ошибки race condition. блокируем  и чтение и запись
	defer l.mtx.Unlock()
	if _, ok := l.tasks[task.Title]; ok { // на момент чтения блокируем так как дальше меняем и чтобы никто не перезаписал map
		return TaskAlreadyExists
	}
	l.tasks[task.Title] = task
	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrNotFound
	}
	return task, nil
}

func (l *List) GetList() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) CompleteTask(title string, status bool) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return ErrNotFound
	}
	task.CompleteTask(status)
	l.tasks[title] = task
	return nil
}

func (l *List) DeleteTask(title string) error { 
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.tasks[title]
	if !ok {
		return ErrNotFound
	}
	delete(l.tasks, title)
	return nil
}

func (l *List) GetNotCompletedTaskList() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.IsCompleted {
			tmp[title] = task
		}
	}
	return tmp
}
