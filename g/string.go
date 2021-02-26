package g

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/johnfrankmorgan/gazebo/errors"
)

var _ Object = &String{}

type String struct {
	Base
	value string
}

func NewString(value string) *String {
	object := &String{value: value}
	object.self = object
	return object
}

func NewStringf(format string, args ...interface{}) *String {
	return NewString(fmt.Sprintf(format, args...))
}

func (m *String) Value() interface{} {
	return m.value
}

func (m *String) String() string {
	return m.value
}

func (m *String) Len() int {
	return len(m.value)
}

// GAZEBO STRING OBJECT METHODS

func (m *String) G_repr() *String {
	return NewStringf("%q", m.value)
}

func (m *String) G_str() *String {
	return NewString(m.value)
}

func (m *String) G_num() *Number {
	value, err := strconv.ParseFloat(m.value, 64)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
	return NewNumber(value)
}

func (m *String) G_bool() *Bool {
	return NewBool(m.value != "")
}

func (m *String) G_len() *Number {
	return NewNumber(float64(m.Len()))
}

func (m *String) G_inverse() Object {
	var (
		buff   bytes.Buffer
		length = m.Len()
	)

	for i := 0; i < length; i++ {
		buff.WriteByte(m.value[length-i-1])
	}

	return NewString(buff.String())
}

func (m *String) G_add(other Object) Object {
	return NewString(m.value + other.G_str().String())
}
