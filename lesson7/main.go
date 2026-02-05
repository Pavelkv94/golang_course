import (
	"fmt"
	"sync"
)

// var number int = 0
//! atomic позволяет выполнять одну непрерывную операцию и предотвращает race condition
// var number atomic.Int64
var slice []int

var mtx sync.Mutex	


func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}

//* mutex
func main7() {

	wg := &sync.WaitGroup{}

	wg.Add(8)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	fmt.Println(len(slice))
}
