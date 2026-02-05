
import (
	"fmt"
	"lesson2/auto"
	"lesson2/methods"
	"lesson2/mine"
	"lesson2/payments"
	"lesson2/user"
	"time"

	"github.com/k0kubun/pp"
)

func main2() {
	bmw := auto.BMW{}
	auto.Ride(bmw)

	paymentMethod := methods.NewPaypalMethod()

	paymentModule := payments.NewPaymentModule(paymentMethod)

	paymentModule.Pay("iphone", 10)
	paymentModule.Pay("Samsung", 15)
	gameId := paymentModule.Pay("Game", 5)
	paymentModule.Cancel(gameId)

	allInfo := paymentModule.AllInfo()

	pp.Println(allInfo)

	// error handling

	user := user.User{
		Name: "Dima", Balance: 10,
	}

	pp.Println("User before payment", user)

	err := user.Pay(15)

	if err != nil {
		pp.Println(err.Error())
	}

	pp.Println("User after payment", user)

	coal := 0

	transferPoint := make(chan int)

	//gorutines
	go mine.Mine(1, transferPoint)
	go mine.Mine(2, transferPoint)
	go mine.Mine(3, transferPoint)

	coal += <-transferPoint
	time.Sleep(1 * time.Second)
	coal += <-transferPoint
	time.Sleep(1 * time.Second)
	coal += <-transferPoint
	fmt.Println("Coals: ", coal)

	go mine.Hello()

	func() {
		fmt.Println("asd")
	}()

	time.Sleep(1 * time.Second)



}
