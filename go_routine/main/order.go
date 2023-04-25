package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	//

	for i := 0; i < 5; i++ {
		i := i
		go func(int) {
			fmt.Println("a" + strconv.Itoa(i))
		}(i)
		go func(int) {
			fmt.Println("b" + strconv.Itoa(i))
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
}
