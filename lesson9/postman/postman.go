package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, transferPoint chan<- string, n int, mail string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон #", n, "мой рабочий день закончен")
			return

		default:
			fmt.Println("Я почтальон #", n, "взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон #", n, "донес письмо до почты ", mail)

			transferPoint <- mail
			fmt.Println("Я почтальон #", n, "отдал письмо на почту ")
		}

	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, mailTransferPoint, i, postmanToMail(i), wg)
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(n int) string {

	ptm := map[int]string{
		1: "Hello",
		2: "How are you?",
		3: "This is bill!",
		4: "Fine 500$",
		5: "test message",
	}

	mail, ok := ptm[n]
	if !ok {
		return "Lottery"
	}
	return mail

}
