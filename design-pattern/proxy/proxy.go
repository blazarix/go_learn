package proxy

import "fmt"

// 代理模式
type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("小汽车行驶中...")
}

// 假如要增加年龄限制又不影响原代码，让一个代理商去代理它，我们直接去代理商那里开车
type Driver struct {
	Age int
}
type CarProxy struct {
	driver  *Driver
	vehicle Vehicle
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{driver, &Car{}}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.vehicle.Drive()
	} else {
		fmt.Println("不符合法定驾驶年龄哦~")
	}
}
