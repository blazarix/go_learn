package factory_method

// 工厂接口，由具体工厂类来实现
type OperatorFactory interface {
	Create() MathOperator
}

// 生产加法计算器的工厂
type PlusOperatorFactory struct{}

func (p *PlusOperatorFactory) Create() MathOperator {
	return &PlusOperator{
		BaseOperator: &BaseOperator{},
	}
}

// 生产乘法计算器的工厂
type MultiOperatorFactory struct{}

func (m *MultiOperatorFactory) Create() MathOperator {
	return &MultiOperator{
		BaseOperator: &BaseOperator{},
	}
}

// 所有计算器的基类
type BaseOperator struct {
	A int
	B int
}

func (o *BaseOperator) SetOperandA(a int) {
	o.A = a
}
func (o *BaseOperator) SetOperandB(b int) {
	o.B = b
}

// 实际的产品--加法计算器
type PlusOperator struct{ *BaseOperator }

func (p *PlusOperator) ComputeResult() int {
	return p.A + p.B
}

// 实际的产品--乘法计算器
type MultiOperator struct{ *BaseOperator }

func (m *MultiOperator) ComputeResult() int {
	return m.A * m.B
}

// 实际产品实现的接口--表示计算器的行为

type MathOperator interface {
	SetOperandA(int)
	SetOperandB(int)
	ComputeResult() int
}
