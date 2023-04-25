package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("指向的chan的地址:%v,本身的地址:%v \n", intChan, &intChan)

	// 写入数据,不能超过其容量
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 92
	// 超过容量会死锁！
	//intChan <- 10
	// 先扔在放不会报错
	<-intChan
	intChan <- 20
	// 查看管道长度和容量
	fmt.Printf("len:%v cap:%v \n", len(intChan), cap(intChan))

	// 取数据,(因为是FIFO，所以会先取出先存放的数据)
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("len:%v cap:%v \n", len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	// 在没有使用协程的情况下，如果没有数据了，还要取数据，会死锁
	//num5 := <-intChan
	//fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)
	fmt.Println("num3=", num3, "num4=", num4)

}
