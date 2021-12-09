package learn1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//查找重复的行
//判断重复输入的行，每次重复输入，map对应的值都会+1，遍历所有map，如果有的key的value大于1就说明输入超过了两次
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	var i int

	for input.Scan() {
		counts[input.Text()]++
		i++
		if i == 3 {
			break
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}
}

//练习，修改为可以打印重复所在的文件名
func Dup2() {
	counts := make(map[string]int)
	fileNameMap := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNameMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts, fileNameMap)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileNameMap[line])
		}
	}
}

//使用ioutil读取文件，一次性读入文件全部内容转换成二进制读到内存中，dup2是一行一行读的。然后用split函数分隔存储
func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data ,err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}
	for line, n := range counts {
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNameMap map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fileNameMap[input.Text()] = f.Name()
		}
	}
}