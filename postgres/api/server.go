package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
}

func NewHttpServer() *HTTPServer {
	return &HTTPServer{
	}
}

func (s *HTTPServer) Start(port string) error {
	app := chi.NewRouter()

	app.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})


	if err := http.ListenAndServe(port, app); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}