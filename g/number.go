package g

import (
	"math"
	"strconv"
)

var _ Object = &Number{}

type Number struct {
	h     ObjectHelper
	value float64
}

func NewNumber(value float64) *Number {
	return &Number{value: value}
}

func (m *Number) Value() interface{} {
	return m.value
}

func (m *Number) Float() float64 {
	return m.value
}

func (m *Number) Int() int64 {
	return int64(m.value)
}

func (m *Number) IsInt() bool {
	return math.Mod(m.value, 1) == 0
}

func (m *Number) CallMethod(name string, args *Args) Object {
	return m.h.CallMethod(m, name, args)
}

func (m *Number) HasAttr(name string) bool {
	return m.h.HasAttr(m, name)
}

func (m *Number) GetAttr(name string) Object {
	return m.h.GetAttr(m, name)
}

func (m *Number) SetAttr(name string, value Object) {
	m.h.SetAttr(m, name, value)
}

func (m *Number) DelAttr(name string) {
	m.h.DelAttr(m, name)
}

// GAZEBO NUMBER OBJECT METHODS

func (m *Number) G_str() *String {
	if m.IsInt() {
		return NewString(strconv.FormatInt(m.Int(), 10))
	}

	return NewString(strconv.FormatFloat(m.value, 'g', -1, 64))
}

func (m *Number) G_num() *Number {
	return NewNumber(m.value)
}

func (m *Number) G_bool() *Bool {
	return NewBool(m.value != 0)
}

func (m *Number) G_not() *Bool {
	return NewBool(!m.G_bool().Bool())
}
