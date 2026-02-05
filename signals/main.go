package signals

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// при получении сигнала от системы, контекст завершается и мы можем завершить горутины без последствий
func main() {
	fmt.Println("PID: ", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM) // системный сигнал, 

	go func() {
		fmt.Println("Waiting for signal...")
		<-ctx.Done()
		fmt.Println("Received signal: ", ctx.Err())
	}()

	
}