package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 获取逻辑cpu数量，并非真实
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum:", cpuNum)
	// 可以自己设置使用多少个cpu
	runtime.GOMAXPROCS(cpuNum - 1)

	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("接收到停止信号")
				return
			default:
				fmt.Println("运行中", time.Now())
			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- true
}
