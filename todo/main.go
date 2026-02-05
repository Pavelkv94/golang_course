
import (
	"fmt"
	"todo/api"
	"todo/todo"
)

func mainTodo() {
	todoList := todo.NewList()
	handlers := api.NewHTTPHandlers(todoList)
	server := api.NewHttpServer(handlers)
	err := server.Start(":5000")
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
	}
}