package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		// внутри контекста у нас канал и если мы канал закроем с помощью отмены контекста то кейс прочитается из канала
		case <-ctx.Done():
			fmt.Println("Foo контекст завершен")

			return
		default:
			fmt.Println("Foo продолжает выполняться")

		}
		time.Sleep(200 * time.Millisecond)
	}
}

func boo(ctx context.Context) {
	for {
		select {
		// внутри контекста у нас канал и если мы канал закроем с помощью отмены контекста то кейс прочитается из канала
		case <-ctx.Done():
			fmt.Println("Boo контекст завершен")

			return
		default:
			fmt.Println("Boo продолжает выполняться")

		}
		time.Sleep(200 * time.Millisecond)
	}
}

//* контекст
func main5() {
	// создали родительский контекст с возможностью отмены
	parentContext, parentCancel := context.WithCancel(context.Background())

	// создали дочерний контекст с возможностью отмены
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(childContext)
	go boo(parentContext)

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(3 * time.Second)
	parentCancel()
}
