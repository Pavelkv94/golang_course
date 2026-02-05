package greeting

import "fmt"

func SayHello() {
	fmt.Println("hello package")
	i := getInt()
	fmt.Println(i)
}