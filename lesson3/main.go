
import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main3() {

	messageCh1 := make(chan Message)
	messageCh2 := make(chan Message)

	go func() {
		for {
			messageCh1 <- Message{
				Author: "Friend 1",
				Text:   "Hello",
			}

			time.Sleep(10000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			messageCh2 <- Message{
				Author: "Friend 2",
				Text:   "How are you?",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-messageCh1:
			fmt.Println("I got message from ", msg1.Author, "He wrote ", msg1.Text)
		case msg2 := <-messageCh2:
			fmt.Println("I got message from ", msg2.Author, "He wrote ", msg2.Text)

		}

	}

}
