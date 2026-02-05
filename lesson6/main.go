
import (
	"fmt"
	"sync"
	"time"
)


func postman(text string, wg *sync.WaitGroup) {
	defer  wg.Done()// завершаем горутину

	for i := 1; i < 5; i++ {
		fmt.Println("I send email with topic: ", text)
		time.Sleep(500 * time.Millisecond)
	}
	
}
// WaitGroup - применяется для горутин которые не возвращают значение
func main6() {

	wg := &sync.WaitGroup{}


	wg.Add(1) // сообщаем сколько горутин запускаем
	go postman("News", wg)
	wg.Add(1)
	go postman("Games", wg)
	wg.Add(1)
	go postman("Auto", wg)

	wg.Wait() // ждем пока выполнятся горутины

}
