package main

import "fmt"

type T1 interface {
	Say()
	Do()
}

type T2 interface {
	Say()
	Do()
}

type S1 struct{}

func (s S1) Say() {}
func (s S1) Do()  {}

func main() {
	// 两个包含相同方法的接口类型是相同的类型
	var t1 T1
	var t2 T2
	fmt.Println(t1 == t2)

}
