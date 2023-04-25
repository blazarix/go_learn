package main

import (
	"fmt"
	"sync"
	"time"
)

// 计算1-200的阶乘，并且把各个数的阶乘放到map中，最后打印出来

var result = make(map[int]int, 10)

// 声明一个全局的互斥锁:mutex互斥
var lock sync.Mutex

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	// 主线程执行完了，goroutine不一定执行完，所以这里等待10s，但是并不知道具体要等多少s，所以这是低级的写法
	time.Sleep(time.Second * 10)

	lock.Lock()
	fmt.Println("result:", result)
	lock.Unlock()
}

func test(n int) {
	r := 1
	for j := 1; j <= n; j++ {
		r *= j
	}
	lock.Lock()
	result[n] = r
	lock.Unlock()
}
