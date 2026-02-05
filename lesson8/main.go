
import (
	"fmt"
	"sync"
)

var likes int = 0

// var mtx sync.RWMutex //! подходит для чтения и записи, но если много чтений то работает медленно(горутины В ОЧЕРЕДИ)

var mtx sync.RWMutex //! эффективен при многих чтениях(проверяет блокировки на запись), если нет записи то позволяет брать данные все горутинам ОДНОВРЕМЕННО

//* go run -race main.go - определение гонки данных 
func setLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100000; i++ {
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100000; i++ {
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}
// rwmutex
func main8() {
	wg := &sync.WaitGroup{}

	for i :=1; i <= 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i :=1; i <= 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()
	fmt.Println(likes)
}
