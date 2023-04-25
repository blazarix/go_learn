package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	// 关闭后不能再写入了
	//close(intChan)
	intChan <- 3

	// 如果不关闭管道进行遍历，会出现deadlock死锁!
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
