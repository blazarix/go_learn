package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 10)

	cat := Cat{"Tom", 12}
	allChan <- cat

	outCat := <-allChan
	// 使用类型断言
	o := outCat.(Cat)

	fmt.Println(o.Name)
}
