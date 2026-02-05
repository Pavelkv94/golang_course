package main

import (
	"context"
	"fmt"
	"os"
	"postgres/api"
	"postgres/db"
	simplesql "postgres/simple_sql"
)

//! для миграций необходимо поставить пакет github.com/golang-migrate/migrate/v4
//! migrate create -ext sql -dir migrations -seq init - шаблон для создания миграции
//! migrate -path migrations -database "postgresql://admin:admin@localhost:5432/gotest?sslmode=disable" up - выполнить миграции
//! migrate -path migrations -database "postgresql://admin:admin@localhost:5432/gotest?sslmode=disable" down - откатить миграции
func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL is not set")
	}

	ctx := context.Background()

	conn, err := db.CreateConnection(ctx, databaseUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	// err = simplesql.CreateTable(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	//! insert
	// err = simplesql.InsertTask(ctx, conn, simplesql.TaskDto{Title: "Покрасить", Description: "купить хлеб в магазине"})
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Task inserted")

	//! update
	// err = simplesql.UpdateTask(ctx, conn, simplesql.UpdateTaskDto{Id: 1, Title: "test title 3", Description: "test description 3"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Task updated")

	//! delete
	// err = simplesql.DeleteTask(ctx, conn, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Task deleted")

	// defer conn.Close(ctx)


	//! select
	tasks, err := simplesql.SelectTasks(ctx, conn)
	if err != nil {
		panic(err)
	}




	fmt.Println(tasks)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}

	server := api.NewHttpServer()
	fmt.Println("Server is running on port: ", port)
	err = server.Start(":" + port)
	if err != nil {
		panic(err)
	}
}