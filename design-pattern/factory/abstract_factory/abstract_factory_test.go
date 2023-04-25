package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	huaWei := &HuaWeiFactory{}
	// 看华为电视
	huaTv := huaWei.CreateTelevision()
	huaTv.Watch()
	// 用华为空调
	huaAir := huaWei.CreateAirConditioner()
	huaAir.SetTemperature(28)

	mi := &MiFactory{}
	// 看小米电视
	miTv := mi.CreateTelevision()
	miTv.Watch()
	// 用小米空调
	miAir := mi.CreateAirConditioner()
	miAir.SetTemperature(27)
}
