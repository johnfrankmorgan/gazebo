package g

import (
	"fmt"
	"strconv"
)

var _ Object = &String{}

type String struct {
	Base
	value string
}

func NewString(value string) *String {
	object := &String{value: value}

	object.SetType(TypeString)
	object.SetSelf(object)

	return object
}

func NewStringf(format string, args ...interface{}) *String {
	return NewString(fmt.Sprintf(format, args...))
}

func (m *String) Value() interface{} {
	return m.value
}

func (m *String) ToBool() *Bool {
	return NewBool(!m.IsEmpty())
}

func (m *String) ToNumber() *Number {
	value, err := strconv.ParseFloat(m.value, 64)
	if err != nil {
		panic(err)
	}

	return NewNumber(value)
}

func (m *String) ToString() *String {
	return NewString(m.value)
}

func (m *String) String() string {
	return m.value
}

func (m *String) Len() int {
	return len(m.value)
}

func (m *String) IsEmpty() bool {
	return m.Len() == 0
}

func (m *String) Limit(length int, extra string) *String {
	if length >= m.Len() {
		return m.ToString()
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	return NewString(m.value[:min(m.Len(), length)] + extra)
}
