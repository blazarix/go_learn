package simple

import "testing"

func TestSimpleFactory(t *testing.T) {
	cn := NewPrinter("cn")
	en := NewPrinter("en")
	println(cn.Print("张三"))
	println(en.Print("张三"))
}
