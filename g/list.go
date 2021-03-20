package g

import (
	"strings"
)

var _ Object = &List{}

type List struct {
	Base
	value []Object
}

func NewList(value []Object) *List {
	object := &List{value: value}

	object.SetType(TypeList)
	object.SetSelf(object)

	return object
}

func NewListSized(size int) *List {
	return NewList(make([]Object, 0, size))
}

func (m *List) Value() interface{} {
	return m.value
}

func (m *List) ToBool() *Bool {
	return NewBool(!m.IsEmpty())
}

func (m *List) ToString() *String {
	var buff strings.Builder

	buff.WriteByte('[')

	for pos, value := range m.value {
		buff.WriteString(value.ToString().String())

		if pos < m.Len()-1 {
			buff.WriteString(", ")
		}
	}

	buff.WriteByte(']')

	return NewString(buff.String())
}

func (m *List) Slice() []Object {
	return m.value
}

func (m *List) Len() int {
	return len(m.value)
}

func (m *List) IsEmpty() bool {
	return m.Len() == 0
}

func (m *List) Get(offset int) Object {
	return m.value[offset]
}

func (m *List) Set(offset int, value Object) {
	m.value[offset] = value
}

func (m *List) Prepend(values ...Object) {
	m.value = append(values, m.value...)
}

func (m *List) Append(values ...Object) {
	m.value = append(m.value, values...)
}
