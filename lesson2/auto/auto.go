package auto

import "fmt"

type BMW struct {
}

func (b BMW) StepOnGas() {
	fmt.Println("BMW is driving")
}

type AUDi struct {
}

func (a AUDi) StepOnGas() {
	fmt.Println("AUDi is driving")
}

// интерфейс реализует авто которое имеет метод StepOnGas (ауди и БМВ)
type Auto interface {
	StepOnGas()
}

func Ride(a Auto) {
	fmt.Println("Auto")
	a.StepOnGas()
}

func main() {
	bmw := BMW{}
	audi := AUDi{}

	Ride(bmw)
	Ride(audi)

}
