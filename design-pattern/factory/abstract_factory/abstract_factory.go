package abstract_factory

import "fmt"

// 创建一个抽象工厂类
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}
type ITelevision interface {
	Watch()
}

type IAirConditioner interface {
	SetTemperature(int)
}

type HuaWeiFactory struct{}

func (h *HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTelevision{}
}

type HuaWeiTelevision struct{}

func (h *HuaWeiTelevision) Watch() {
	fmt.Println("欢迎使用华为电视")
}

func (h *HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{}
}

type HuaWeiAirConditioner struct{}

func (h *HuaWeiAirConditioner) SetTemperature(t int) {
	fmt.Printf("欢迎使用华为空调，调节温度为 %d 度\n", t)
}

type MiFactory struct{}

func (h *MiFactory) CreateTelevision() ITelevision {
	return &MiTelevision{}
}

type MiTelevision struct{}

func (h *MiTelevision) Watch() {
	fmt.Println("欢迎使用小米电视")
}

func (h *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}

type MiAirConditioner struct{}

func (h *MiAirConditioner) SetTemperature(t int) {
	fmt.Printf("欢迎使用小米空调，调节温度为 %d 度\n", t)
}
