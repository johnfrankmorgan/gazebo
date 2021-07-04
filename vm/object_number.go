package vm

var _ Object = Number{}

type Number struct {
	LazyAttributes
	value float64
}

func NewNumber(value float64) Number {
	return Number{value: value}
}

func (m Number) Type() Type {
	return Types.Number
}

func (m Number) Value() interface{} {
	return m.value
}

func (m Number) Float() float64 {
	return m.value
}

func (m Number) Int() int {
	return int(m.value)
}

func (m Number) UInt64() uint64 {
	return uint64(m.value)
}
