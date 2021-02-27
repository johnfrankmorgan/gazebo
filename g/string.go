package g

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/johnfrankmorgan/gazebo/errors"
)

var _ Object = &String{}

type String struct {
	Base
	value string
}

func NewString(value string) *String {
	object := &String{value: value}
	object.SetSelf(object)
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

// GAZEBO STRING OBJECT PROTOCOLS

func (m *String) G_repr() *String {
	return NewStringf("%q", m.value)
}

func (m *String) G_str() *String {
	return NewString(m.value)
}

func (m *String) G_num() *Number {
	value, err := strconv.ParseFloat(m.value, 64)
	errors.ErrRuntime.ExpectNilError(err)
	return NewNumber(value)
}

func (m *String) G_bool() *Bool {
	return NewBool(m.value != "")
}

func (m *String) G_len() *Number {
	return NewNumberFromInt(m.Len())
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

func (m *String) G_contains(value Object) *Bool {
	return NewBool(strings.Contains(m.value, value.G_str().String()))
}

// GAZEBO STRING OBJECT METHODS

func (m *String) G_lower() *String {
	return NewString(strings.ToLower(m.value))
}

func (m *String) G_upper() *String {
	return NewString(strings.ToUpper(m.value))
}

func (m *String) G_find(value Object) *Number {
	index := strings.Index(m.value, value.G_str().String())
	return NewNumberFromInt(index)
}

func (m *String) G_replace(find, replace Object) *String {
	return NewString(strings.ReplaceAll(
		m.value,
		find.G_str().String(),
		replace.G_str().String(),
	))
}

func (m *String) G_empty() *Bool {
	return NewBool(m.Len() == 0)
}

func (m *String) G_from(index Object) *String {
	return NewString(m.value[index.G_num().Int():])
}

func (m *String) G_until(index Object) *String {
	return NewString(m.value[:index.G_num().Int()])
}

func (m *String) G_slice(start, end Object) *String {
	return NewString(m.value[start.G_num().Int():end.G_num().Int()])
}

func (m *String) G_numeric() *Bool {
	regex := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	return NewBool(regex.MatchString(m.value))
}
