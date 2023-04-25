package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// A:
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Println("存储了：", Balance())
	}()
	// B:
	go func() {
		defer wg.Done()
		Deposit(100)
	}()

	wg.Wait()

	for i := 0; i < 400; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("i=", i)
	}
}
