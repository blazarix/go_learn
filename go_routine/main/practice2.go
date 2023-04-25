package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 开一个协程生成1000条数据，存放到文件中
// 完成1000个数据写入后，让sort协程从文件中读取1000个数据，并完成排序，重新写入到另一个文件
func main() {
	successWrite := make(chan bool, 1)
	successRead := make(chan bool, 1)
	go writeData(successWrite)
	go readData(successWrite, successRead)
	for {
		if _, ok := <-successRead; !ok {
			break
		}
	}

}

func readData(successWrite, successRead chan bool) {

	<-successWrite

	nums, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	var arr []int
	arrStr := strings.Split(string(nums), "\n")
	for _, s := range arrStr {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}

	// 排序
	sort.Ints(arr)
	// 写入另一个文件
	s := bytes.Buffer{}
	for _, num := range arr {
		s.WriteString(strconv.Itoa(num) + "\n")
	}
	// 创建并写入数据到文件
	err = ioutil.WriteFile("sort.txt", s.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	successRead <- true
	close(successRead)
}

func writeData(successWrite chan bool) {
	s := bytes.Buffer{}
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 1000; i++ {
		num := rand.Intn(2000)
		fmt.Println("生成随机数", num)
		s.WriteString(strconv.Itoa(num) + "\n")
	}
	// 创建并写入数据到文件
	err := ioutil.WriteFile("test.txt", s.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	successWrite <- true
	close(successWrite)
}
