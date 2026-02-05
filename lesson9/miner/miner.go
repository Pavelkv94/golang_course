package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context, transferPoint chan<- int, n int, power int, wg *sync.WaitGroup) { //chan<- значит канал только для записи
	defer wg.Done()

	for {

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер #", n, "Мой нрабочий день закончен")
			return
		default:
			fmt.Println("Я шахтер #", n, "начал добывать уголь")
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтер #", n, "добыл уголь в количестве ", power)

			transferPoint <- power
			fmt.Println("Я шахтер #", n, "передал уголь в количестве ", power)
		}

	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int { //  <-chan только для чтения
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, coalTransferPoint, i, i*10, wg)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
