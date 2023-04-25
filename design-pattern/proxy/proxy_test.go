package proxy

import "testing"

func TestProxy(t *testing.T) {
	car := NewCarProxy(&Driver{Age: 18})
	car.Drive()
	car2 := NewCarProxy(&Driver{Age: 12})
	car2.Drive()
}
