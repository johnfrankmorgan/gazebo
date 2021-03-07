package g

import (
	"bytes"
)

var _ Object = &List{}

type List struct {
	Base
	value []Object
}

func NewList(value []Object) *List {
	object := &List{value: value}
	object.SetSelf(object)
	return object
}

func NewListSized(size int) *List {
	return NewList(make([]Object, size))
}

func (m *List) Value() interface{} {
	return m.value
}

func (m *List) Append(objects ...Object) *List {
	m.value = append(m.value, objects...)
	return m
}

func (m *List) Prepend(objects ...Object) *List {
	old := m.value
	m.value = objects
	m.Append(old...)
	return m
}

func (m *List) All() []Object {
	return m.value
}

func (m *List) From(index int) *List {
	return NewList(m.value[index:])
}

func (m *List) Until(index int) *List {
	return NewList(m.value[:index])
}

func (m *List) Slice(start, end int) *List {
	return NewList(m.value[start:end])
}

func (m *List) Empty() bool {
	return m.Len() == 0
}

func (m *List) Len() int {
	return len(m.value)
}

func (m *List) Has(index int) bool {
	return index >= 0 && index < m.Len()
}

func (m *List) Get(index int) Object {
	return m.value[index]
}

func (m *List) Set(index int, value Object) {
	m.value[index] = value
}

func (m *List) Contains(value Object) bool {
	for _, val := range m.value {
		if val.G_eq(value).Bool() {
			return true
		}
	}

	return false
}

// GAZEBO LIST OBJECT PROTOCOLS

func (m *List) G_str() *String {
	var (
		buff   bytes.Buffer
		length = m.Len()
	)

	buff.WriteByte('[')

	for i, value := range m.All() {
		buff.WriteString(value.G_repr().String())

		if i < length-1 {
			buff.WriteString(", ")
		}
	}

	buff.WriteByte(']')

	return NewString(buff.String())
}

func (m *List) G_bool() *Bool {
	return NewBool(!m.Empty())
}

func (m *List) G_len() *Number {
	return NewNumberFromInt(m.Len())
}

func (m *List) G_inverse() Object {
	var (
		length = m.Len()
		list   = NewListSized(length)
	)

	for i, object := range m.value {
		list.Set(length-i-1, object)
	}

	return list
}

func (m *List) G_getattr(name *String) Object {
	if name.String() == "length" {
		return m.G_len()
	}

	return m.Base.G_getattr(name)
}
func (m *List) G_contains(value Object) *Bool {
	return NewBool(m.Contains(value))
}

// GAZEBO LIST OBJECT METHODS

func (m *List) G_has(index Object) *Bool {
	return NewBool(m.Has(index.G_num().Int()))
}

func (m *List) G_get(index Object) Object {
	return m.Get(index.G_num().Int())
}

func (m *List) G_set(index, value Object) {
	m.Set(index.G_num().Int(), value)
}

func (m *List) G_append(values ...Object) *List {
	return m.Append(values...)
}

func (m *List) G_prepend(values ...Object) *List {
	return m.Prepend(values...)
}

func (m *List) G_empty() *Bool {
	return NewBool(m.Empty())
}

func (m *List) G_from(index Object) *List {
	return m.From(index.G_num().Int())
}

func (m *List) G_until(index Object) *List {
	return m.Until(index.G_num().Int())
}

func (m *List) G_slice(start, end Object) *List {
	return m.Slice(start.G_num().Int(), end.G_num().Int())
}

func (m *List) G_all() *Bool {
	for _, obj := range m.All() {
		if obj.G_not().Bool() {
			return NewBool(false)
		}
	}

	return NewBool(true)
}

func (m *List) G_any() *Bool {
	for _, obj := range m.All() {
		if obj.G_bool().Bool() {
			return NewBool(true)
		}
	}

	return NewBool(false)
}

func (m *List) G_filter(cb Object) *List {
	list := NewList(nil)

	for i, obj := range m.All() {
		if cb.G_invoke(NewVarArgs(obj, NewNumberFromInt(i))).G_bool().Bool() {
			list.Append(obj)
		}
	}

	return list
}

func (m *List) G_map(cb Object) *List {
	list := NewListSized(m.Len())

	for i, obj := range m.All() {
		list.Set(i, cb.G_invoke(NewVarArgs(obj, NewNumberFromInt(i))))
	}

	return list
}

func (m *List) G_each(cb Object) {
	for i, obj := range m.All() {
		cb.G_invoke(NewVarArgs(obj, NewNumberFromInt(i)))
	}
}
