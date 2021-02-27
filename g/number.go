package g

import (
	"math"
	"strconv"
)

var _ Object = &Number{}

type Number struct {
	Base
	value float64
}

func NewNumber(value float64) *Number {
	object := &Number{value: value}
	object.SetSelf(object)
	return object
}

func NewNumberFromInt(value int) *Number {
	return NewNumber(float64(value))
}

func NewNumberFromInt64(value int64) *Number {
	return NewNumber(float64(value))
}

func (m *Number) Value() interface{} {
	return m.value
}

func (m *Number) IsInt() bool {
	return math.Mod(m.value, 1) == 0
}

func (m *Number) Float() float64 {
	return m.value
}

func (m *Number) Int() int {
	return int(m.value)
}

func (m *Number) Int64() int64 {
	return int64(m.value)
}

// GAZEBO NUMBER OBJECT PROTOCOLS

func (m *Number) G_repr() *String {
	return m.G_str()
}

func (m *Number) G_str() *String {
	if m.IsInt() {
		return NewString(strconv.FormatInt(m.Int64(), 10))
	}

	return NewString(strconv.FormatFloat(m.value, 'g', -1, 64))
}

func (m *Number) G_num() *Number {
	return NewNumber(m.value)
}

func (m *Number) G_bool() *Bool {
	return NewBool(m.value != 0)
}

func (m *Number) G_inverse() Object {
	return NewNumber(-m.value)
}

func (m *Number) G_add(other Object) Object {
	return NewNumber(m.value + other.G_num().value)
}

func (m *Number) G_sub(other Object) Object {
	return NewNumber(m.value - other.G_num().value)
}

func (m *Number) G_mul(other Object) Object {
	return NewNumber(m.value * other.G_num().value)
}

func (m *Number) G_div(other Object) Object {
	return NewNumber(m.value / other.G_num().value)
}

func (m *Number) G_gt(other Object) *Bool {
	return NewBool(m.value > other.G_num().Float())
}

func (m *Number) G_gte(other Object) *Bool {
	return NewBool(m.value >= other.G_num().Float())
}

func (m *Number) G_lt(other Object) *Bool {
	return NewBool(m.value < other.G_num().Float())
}

func (m *Number) G_lte(other Object) *Bool {
	return NewBool(m.value <= other.G_num().Float())
}
