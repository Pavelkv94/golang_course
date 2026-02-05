package mine

import (
	"fmt"
	"time"
)

func Mine(i int, ch chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Добываю из шахты: ", i)

	ch <- 10
	fmt.Println("Отдал уголь в пункт из шахты: ", i)

}

func Hello() {
	for {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)

	}

}