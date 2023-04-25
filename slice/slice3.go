package main

import "fmt"

func main() {
	//var a []int
	a := make([]int, 1)
	a[0] = 1
	fmt.Println(len(a), cap(a))
}
