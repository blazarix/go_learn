package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
	Age int
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Student{Age: 12}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
