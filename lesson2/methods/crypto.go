package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct {
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Payment with cryptocurrency: ", usd, "USDT")

	id := rand.Int();

	return id
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Payment with cryptocurrency canceled: operation #", id)
}

func NewCryptoMethod() Crypto {
	return Crypto{}
}
