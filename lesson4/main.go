

import (
	"fmt"
	"math/rand"
	"time"
)

//* закрытие каналов
func main4() {

	transferPoint := make(chan int)

	go func() {
		iterations := 3 + rand.Intn(4)
		fmt.Println(iterations)

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(300 * time.Millisecond)
		}
		close(transferPoint)

	}()

	coal := 0

	// for {
		
	// 	v, ok := <-transferPoint
	// 	if !ok { //! канал закрыт
	// 		break
	// 	}

	// 	coal += v
	// 	fmt.Println(coal)
	// }

	//! то же самое что выше только без ok. 
	for v := range transferPoint {

		coal += v
		fmt.Println(coal)
	}

}
