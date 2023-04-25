package main

import "fmt"

func main() {
	e := []int32{1, 2, 3}
	a := []int32{1, 2, 3, 4}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, a...)
	fmt.Println("cap of e after:", cap(e))
}
