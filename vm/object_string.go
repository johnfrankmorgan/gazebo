package vm

import "fmt"

var _ Object = &String{}

type String struct {
	LazyAttributes
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func NewStringf(format string, args ...interface{}) *String {
	return NewString(fmt.Sprintf(format, args...))
}

func (m *String) Type() Type {
	return Types.String
}

func (m *String) Value() interface{} {
	return m.value
}

func (m *String) String() string {
	return m.value
}
