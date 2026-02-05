package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct {
}

func (p Paypal) Pay(usd int) int {
	fmt.Println("Payment with Paypal: ", usd, "$")

	id := rand.Int();

	return id
}

func (p Paypal) Cancel(id int) {
	fmt.Println("Payment with Paypal canceled: operation #", id)
}

func NewPaypalMethod() Paypal {
	return Paypal{}
}
