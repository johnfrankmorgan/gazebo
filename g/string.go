package g

import (
	"fmt"
	"strconv"

	"github.com/johnfrankmorgan/gazebo/errors"
)

var _ Object = &String{}

type String struct {
	Partial
	h     ObjectHelper
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

func (m *String) CallMethod(name string, args *Args) Object {
	return m.h.CallMethod(m, name, args)
}

func (m *String) HasAttr(name string) bool {
	return m.h.HasAttr(m, name)
}

func (m *String) GetAttr(name string) Object {
	return m.h.GetAttr(m, name)
}

func (m *String) SetAttr(name string, value Object) {
	m.h.SetAttr(m, name, value)
}

func (m *String) DelAttr(name string) {
	m.h.DelAttr(m, name)
}

// GAZEBO STRING OBJECT METHODS

func (m *String) G_str() *String {
	return &String{value: m.value}
}

func (m *String) G_num() *Number {
	value, err := strconv.ParseFloat(m.value, 64)
	errors.ErrRuntime.ExpectNil(err, "%v", err)
	return NewNumber(value)
}

func (m *String) G_bool() *Bool {
	return NewBool(m.value != "")
}

func (m *String) G_not() *Bool {
	return NewBool(!m.G_bool().Bool())
}
