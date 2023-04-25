package main

import (
	"fmt"
	"sync"
)

var rwMutex sync.RWMutex
var data []int

func readData() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	// 读取数据
	fmt.Println("data:", data)
}

func writeData() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	data = append(data, 2)
}

func main() {

}
