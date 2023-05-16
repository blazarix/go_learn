package main

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}

	arr2 := arr1

	arr2 = append(arr2[:2], arr2[3:]...)

	//fmt.Println(arr1)
	//fmt.Println(arr2)
}
