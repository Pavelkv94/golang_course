package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var money atomic.Int64
var bank atomic.Int64
var mtx sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	str := "hello world"
	b := []byte(str)
	time.Sleep(3 * time.Second)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Sended")
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		msg := "fail to get request body: " + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail")
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	amount, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // status code
		msg := "fail to get request body: " + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail")
		}
		return
	}

	mtx.Lock() //! примееняем мьютекс потому что в разных хэндлерахз используем одни и те же глобальные переменные, атомик действует на переменную а не на целые операции - можно заменить на инт
	if money.Load()-int64(amount) >= 0 {
		time.Sleep(3 * time.Second)
		money.Add(int64(-amount))
		fmt.Println("You money after paid: ", money.Load())
		w.Write([]byte("SUCCESS"))
	} else {
		fmt.Println("You have not enough money")

	}
	mtx.Unlock() //! примееняем мьютекс потому что в разных хэндлерахз используем одни и те же глобальные переменные, атомик действует на переменную а не на целые операции - можно заменить на инт

}

func saveMoneyHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("fail to get request body: ", err.Error())
		return 
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		fmt.Println("fail to get request body: ", err.Error())
		return
	}

	mtx.Lock() //! примееняем мьютекс потому что в разных хэндлерахз используем одни и те же глобальные переменные, атомик действует на переменную а не на целые операции - можно заменить на инт
	if money.Load() >= int64(saveAmount) {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("Money: ", money.Load())
		fmt.Println("Bank: ", bank.Load())
	} else {
		fmt.Println("You have not enough money for saving")

	}
	mtx.Unlock() //! примееняем мьютекс потому что в разных хэндлерахз используем одни и те же глобальные переменные, атомик действует на переменную а не на целые операции - можно заменить на инт

}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Cancel payment"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Canceled")
	}
}

func main_1() {
	money.Add(1000)

	http.HandleFunc("/", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveMoneyHandler)
	http.HandleFunc("/pay/cancel", cancelHandler)
	fmt.Println("Server is Running on 5000")

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("Server Error: ", err.Error())
	}

}
