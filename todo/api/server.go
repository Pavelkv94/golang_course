package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHttpServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) Start(port string) error {
	app := chi.NewRouter()

	app.Post("/tasks", s.httpHandlers.AddTaskHandler)
	app.Get("/tasks", s.httpHandlers.GetTasksHandler)
	app.Get("/tasks/{title}", s.httpHandlers.GetTaskHandler)
	app.Get("/tasks?not_completed=true", s.httpHandlers.GetNotCompletedTasksHandler)
	app.Patch("/tasks/{title}", s.httpHandlers.CompleteTaskHandler)
	app.Delete("/tasks/{title}", s.httpHandlers.DeleteTaskHandler)

	if err := http.ListenAndServe(port, app); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}