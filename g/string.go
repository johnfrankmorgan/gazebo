package g

import "fmt"

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
	return NewBool(m.Len() > 0)
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
