package objects

import (
	"bytes"
	"fmt"
)

var _ Object = &String{}

type String struct {
	buffer bytes.Buffer
}

func NewString(value string) *String {
	var s String

	s.buffer.WriteString(value)

	return &s
}

func NewStringf(format string, args ...interface{}) *String {
	var s String

	fmt.Fprintf(&s.buffer, format, args...)

	return &s
}

func (m *String) Value() string {
	return m.buffer.String()
}

/* Object methods */

func (m *String) GoVal() interface{} {
	return m.Value()
}

func (m *String) Hash() interface{} {
	return m.Value()
}

func (m *String) Type() Type {
	return TypeString
}

func (m *String) ToString() *String {
	return m
}
