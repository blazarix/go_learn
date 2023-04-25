package factory_method

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	plusFactory := new(PlusOperatorFactory)
	plus := plusFactory.Create()
	plus.SetOperandA(1)
	plus.SetOperandB(2)
	println(plus.ComputeResult())

	multiFactory := new(MultiOperatorFactory)
	multi := multiFactory.Create()
	multi.SetOperandA(6)
	multi.SetOperandB(5)
	println(multi.ComputeResult())
}
