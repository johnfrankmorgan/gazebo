package g

import "math"

var _ Object = &Number{}

type Number struct {
	Base
	value float64
}

func NewNumber(value float64) *Number {
	object := &Number{value: value}

	object.SetType(TypeNumber)
	object.SetSelf(object)

	return object
}

func NewNumberFromInt(value int) *Number {
	return NewNumber(float64(value))
}

func (m *Number) Value() interface{} {
	if m.IsInt() {
		return m.Int()
	}

	return m.Float()
}

func (m *Number) IsInt() bool {
	return math.Mod(m.value, 1) == 0
}

func (m *Number) Int() int {
	return int(m.value)
}

func (m *Number) Float() float64 {
	return m.value
}
