package learn1

import (
	"fmt"
	"os"
	"strings"
)

//命令行参数
func Echo() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "-"
	}
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
	fmt.Println(strings.Join(os.Args, "+"))

	fmt.Println(s)
}

