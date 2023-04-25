package main

import "fmt"

// 无缓冲通道，主线程阻塞，导致执行不到go。
func main() {
	c := make(chan int)
	for i := 0; i <= 10; i++ {
		c <- i
	}

	go func() {
		<-c
		for item := range c {
			fmt.Println(item)
		}
	}()
}
