package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"study/greeting"
	"study/user"

	"github.com/k0kubun/pp"
)

func mainLesson1() {
	user1 := user.NewUser(
		"Dima",
		64,
		"6465465465",
		true,
		5.0)

	fmt.Println("User before methods call", user1)
	user1.Greeting()
	user1.ChangeName("Arkadiy")
	user1.ChangeAge(25)
	user1.ChangePhoneNumber("123123132")
	user1.CloseAccount()
	user1.RatingUp()
	fmt.Println("User after methods call", user1)

	greeting.SayHello()
	greeting.SayBad()

	arr := [5]int{1, 2, 3, 4, 5}
	pp.Println("arr, ----- ", cap(arr)) // capacity

	pp.Println(user1)

	pp.Println(arr[0])

	arr[0] += 5

	pp.Println(arr[4])

	userArr := []user.User{
		user.NewUser(
			"Dima",
			64,
			"6465465465",
			true,
			5.0),
		user.NewUser(
			"Dima",
			64,
			"6465465465",
			true,
			5.0),
	}

	pp.Println(len(userArr))

	//! создает копии внутри, лучше для read операций
	for i, v := range userArr {
		pp.Println(i, v)
		// userArr[i].name = "asd" //! так можно менять исходный обьект
	}

	//! задаем то же имя чтобы не копировать массив а перезаписывать
	userArr = append(userArr, user.NewUser(
		"ANNA",
		64,
		"111111111",
		true,
		2.0))

		pp.Println(userArr)
		pp.Println(len(userArr)) // length
		pp.Println(cap(userArr)) // capacity

		intArr := make([]int, 0) // создаем слайс с длиной 0 который будет увеличиватся
		pp.Println(len(intArr))

		intArr2 := make([]int, 0, 5) // создаем слайс с длиной 5 и без элементов
		pp.Println(len(intArr2))


		weather := map[string]int{
			"11": +2,
			"12": -2,
			"13": +1,
			"14": 11,
		}

		pp.Println(weather["12"])

		c, ok := weather["12"] // c - значение, ok - булево выражение, показывает знеачение заданное или по умолчанию(для map)

		pp.Println(c, ok)

		newmap := make(map[int]int, 10) // создаем map и предвыделяем ему 10 длинну в памяти
		pp.Println(newmap)


		//! Пользовательский ввод
		scanner := bufio.NewScanner(os.Stdin)

		pp.Println("Enter command:")

		inputOk := scanner.Scan(); // то что ввод прочитан успешно или нет

		if(!inputOk) {
			pp.Println("Failed input")
			return
		}

		// if inputOk2 := scanner.Scan(); !inputOk2 { // то же самое что выше только другая запись
		// 	pp.Println("Failed input")
		// 	return
		// }

		text := scanner.Text()

		splittedValues := strings.Fields(text) // разбить строку по пробелам

		fmt.Println("text: ", text)
		fmt.Println("splittedValues: ", splittedValues)

		

}
