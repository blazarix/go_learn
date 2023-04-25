package simple

import "fmt"

// 简单工厂模式

type Printer interface {
	Print(name string) string
}

// 向客户端提供需要的语种的打印机

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(cnPrinter)
	case "en":
		return new(enPrinter)
	default:
		return new(cnPrinter)
	}
}

type cnPrinter struct{}

func (c *cnPrinter) Print(name string) string {
	return fmt.Sprintf("你好:%s", name)
}

type enPrinter struct{}

func (c *enPrinter) Print(name string) string {
	return fmt.Sprintf("hello:%s", name)
}
