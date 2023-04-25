package main

import "fmt"

type Res struct {
	Num int
	Sum int
}

// 启动一个协程，将1-20000的数放到一个numChan中，启动8个协程，从numChan中取出数，并计算1+。。。。+n的值存放到resChan。遍历resChan

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan Res, 2000)

	go write(numChan)

	for i := 0; i < 8; i++ {
		go read(numChan, resChan)
	}

	for {
		if len(resChan) == 2000 {
			close(resChan)
			break
		}
	}

	for res := range resChan {
		fmt.Printf("res[%v] = %v \n", res.Num, res.Sum)
	}
}

func write(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
		fmt.Println("写入数据", i)
	}
	close(numChan)
}

func read(numChan chan int, resChan chan Res) {
	for {
		// 对于已经关闭的管道，怎么取都不会报错，可以用ok来判断是否取成功
		num, ok := <-numChan
		if !ok {
			break
		}
		sum := (num * (num + 1)) / 2
		res := Res{
			Num: num,
			Sum: sum,
		}
		fmt.Println("读取到数据", res)
		resChan <- res
	}
}
