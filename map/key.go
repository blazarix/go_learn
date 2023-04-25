package main

import "fmt"

type T1 struct {
	Age  int
	Name string
}

func (t T1) Say() {
	fmt.Println("T1")
}

type T2 struct {
	Age  *int
	Name string
}

type T3 struct {
	Age  []int
	Name string
}

type T4 interface {
	Say()
}

type T5 struct {
	Age  []int
	Name string
}

func main() {
	// 含有内建类型的结构体，可以作为map的key
	t1 := make(map[T1]int)
	a := T1{Age: 1, Name: "a"}
	t1[a] = 1
	fmt.Println(t1)

	// 含有指针类型的结构体，可以作为map的key
	t2 := make(map[T2]int)
	b := T2{Age: &a.Age, Name: "b"}
	t2[b] = 1
	fmt.Println(t2)

	// 数组类型可以作为map的key
	t7 := make(map[[3]int]int)
	j := [3]int{1, 2, 3}
	t7[j] = 1
	fmt.Println(t7)
	// 因为数组可以用==比较
	k := [3]int{1, 2, 3}
	l := [3]int{1, 2, 3}
	fmt.Println(k == l)

	// 含有切片类型的结构体，不能作为map的key
	//t3 := make(map[T3]int)
	//c := T3{Age: []int{1, 2, 3}, Name: "c"}
	//t3[c] = 1
	//fmt.Println(t3)
	// 切片不能作为map的key，因为无法用==比较
	//var list1 []int
	//var list2 []int
	//list1 = append(list1, 1)
	//list2 = append(list2, 1)
	//fmt.Println(list1 == list2)

	// 含有切片的结构体不能用==比较
	//d := T3{Age: []int{1, 2, 3}, Name: "d"}
	//e := T3{Age: []int{1, 2, 3}, Name: "e"}
	//fmt.Println(d == e)

	// 接口类型可以作为map的key
	t4 := make(map[T4]int)
	f := T1{Age: 1, Name: "f"}
	t4[f] = 1
	fmt.Println(t4)

	// 接口可以使用==比较
	var g, h T4
	g = T1{Age: 1, Name: "g"}
	h = T1{Age: 1, Name: "h"}
	fmt.Println(g == h)

	// 函数类型不可以作为map的key
	//t5 := make(map[func() int]int)

	// 因为函数类型不可以用==比较
	//func1 := func() int {
	//	return 1
	//}
	//func2 := func() int {
	//	return 1
	//}
	//fmt.Println(func1 == func2)

	// 指针类型可以作为map的key
	t6 := make(map[*T1]int)
	i := T1{Age: 1, Name: "i"}
	t6[&i] = 1
	fmt.Println(t6)

	// 含有指针类型的结构体，可以用==比较
	var age1 = 1
	var age2 = 1
	tt1 := T2{Age: &age1, Name: "b"}
	tt2 := T2{Age: &age2, Name: "b"}
	fmt.Println(tt1 == tt2)
}
