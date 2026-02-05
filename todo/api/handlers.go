package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo/todo"

	"github.com/go-chi/chi/v5"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	// десериализуем тело запроса в структуру
	var input AddTaskDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		http.Error(w, errorDto.ToString(), http.StatusBadRequest)

		return
	}

	// валидируем структуру
	err = input.ValidateForAddTask()
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		http.Error(w, errorDto.ToString(), http.StatusBadRequest)
		return
	}

	// создаем новую задачу
	newTodoTask := todo.NewTask(input.Title, input.Description)
	// добавляем задачу в список
	err = h.todoList.AddTask(newTodoTask)
	// если выпала ошибка которую мы предусмотрели, то возвращаем ошибку
	if err != nil {
		if errors.Is(err, todo.TaskAlreadyExists) {
			errorDto := ErrorDto{
				Message:   err.Error(),
				Timestamp: time.Now(),
			}
			http.Error(w, errorDto.ToString(), http.StatusConflict) //! если выпала ошибка которую мы предусмотрели, то возвращаем ошибку
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError) //! иначе возвращаем серверную ошибку

		}

		return
	}

	// сериализуем задачу в json
	jsonResponse, err := json.MarshalIndent(newTodoTask, "", "  ")
	if err != nil {
		panic(err)
	}

	// отправляем задачу в ответе
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}

}

func (h *HTTPHandlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	// достаем параметр из url
	param := chi.URLParam(r, "title")

	// получаем задачу из списка
	task, err := h.todoList.GetTask(param)
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		// если выпала ошибка которую мы предусмотрели, то возвращаем ошибку
		if errors.Is(err, todo.ErrNotFound) {
			http.Error(w, errorDto.ToString(), http.StatusNotFound)
		} else { // иначе возвращаем серверную ошибку
			http.Error(w, errorDto.ToString(), http.StatusInternalServerError)
		}
		return

	}

	// сериализуем задачу в json
	jsonResponse, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		panic(err)
	}

	// отправляем задачу в ответе
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error()) // выводим ошибку в консоль потому что она не будет отправлена в ответе
		return
	}

}

func (h *HTTPHandlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.GetList()
	jsonResponse, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}

}

func (h *HTTPHandlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input CompleteTaskDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		http.Error(w, errorDto.ToString(), http.StatusBadRequest)
		return
	}

	err = input.ValidateForCompleteTask()
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		http.Error(w, errorDto.ToString(), http.StatusBadRequest)
		return
	}

	title := chi.URLParam(r, "title")
	err = h.todoList.CompleteTask(title, *input.Completed)
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		if errors.Is(err, todo.ErrNotFound) {
			http.Error(w, errorDto.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errorDto.ToString(), http.StatusInternalServerError)
		}
		return
	}
}

func (h *HTTPHandlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	err := h.todoList.DeleteTask(title)
	if err != nil {
		errorDto := ErrorDto{
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		if errors.Is(err, todo.ErrNotFound) {
			http.Error(w, errorDto.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errorDto.ToString(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Task deleted"))
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}

}

func (h *HTTPHandlers) GetNotCompletedTasksHandler(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.GetNotCompletedTaskList()
	jsonResponse, err := json.MarshalIndent(uncompletedTasks, "", "  ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		return
	}
}
