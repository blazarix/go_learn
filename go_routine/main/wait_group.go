package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(msg string) {
	wg.Add(1)
	defer wg.Done()
	fmt.Printf("worker do %s\n", msg)
}

func main() {
	go worker("task 1")
	go worker("task 2")
	go worker("task 3")
	fmt.Println("waiting")
	wg.Wait()
	fmt.Println("main exit")
}
