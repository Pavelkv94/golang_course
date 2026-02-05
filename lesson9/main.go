package main

import (
	"context"
	"fmt"
	"lesson9/miner"
	"lesson9/postman"
	"time"
)

// concurrency final example - полный пиздос
func main9() {
	var coal int
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerContext, 2)
	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	isCoalClosed := false
	isMailClosed := false

	go func(){
		time.Sleep(3 * time.Second)
		minerCancel()
	}()
	go func(){
		time.Sleep(6 * time.Second)
		postmanCancel()
	}()

	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coalTransferPoint:
			if(!ok) {
				isCoalClosed = true
				continue
			}
			coal += c

		case m, ok := <-mailTransferPoint:
			if(!ok) {
				isMailClosed = true
				continue
			}
			mails = append(mails, m)
		}
	}

	fmt.Println("Суммарное количество угля: ", coal)
	fmt.Println("Суммарное количество писем: ", len(mails))

}
