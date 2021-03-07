package g

import (
	"bytes"
)

var _ Object = &Map{}

type Map struct {
	Base
	values map[uint64]Object
	keys   map[uint64]Object
}

func NewMap() *Map {
	object := &Map{
		keys:   make(map[uint64]Object),
		values: make(map[uint64]Object),
	}
	object.SetSelf(object)
	return object
}

func (m *Map) Value() interface{} {
	return m.values
}

func (m *Map) Empty() bool {
	return m.Len() == 0
}

func (m *Map) Len() int {
	return len(m.values)
}

func (m *Map) Has(key Object) bool {
	_, ok := m.values[key.Hash()]
	return ok
}

func (m *Map) Get(key Object) Object {
	if m.Has(key) {
		return m.values[key.Hash()]
	}

	return NewNil()
}

func (m *Map) Set(key Object, value Object) {
	hash := key.Hash()
	m.keys[hash] = key
	m.values[hash] = value
}

func (m *Map) Del(key Object) {
	hash := key.Hash()

	delete(m.keys, hash)
	delete(m.values, hash)
}

func (m *Map) Contains(key Object) bool {
	return m.Has(key)
}

func (m *Map) HasValue(value Object) bool {
	for _, val := range m.values {
		if val.G_eq(value).Bool() {
			return true
		}
	}

	return false
}

func (m *Map) Keys() map[uint64]Object {
	return m.keys
}

func (m *Map) Values() map[uint64]Object {
	return m.values
}

// GAZEBO LIST OBJECT PROTOCOLS

func (m *Map) G_str() *String {
	var (
		idx    int
		buff   bytes.Buffer
		length = m.Len()
	)

	buff.WriteByte('{')

	for hash, key := range m.keys {
		value := m.values[hash]
		buff.WriteString(key.G_repr().String())
		buff.WriteString(": ")
		buff.WriteString(value.G_repr().String())

		if idx < length-1 {
			buff.WriteString(", ")
		}

		idx++
	}

	buff.WriteByte('}')

	return NewString(buff.String())
}

func (m *Map) G_bool() *Bool {
	return NewBool(!m.Empty())
}

func (m *Map) G_len() *Number {
	return NewNumberFromInt(m.Len())
}

func (m *Map) G_contains(key Object) *Bool {
	return NewBool(m.Contains(key))
}

// GAZEBO MAP OBJECT METHODS

func (m *Map) G_has(key Object) *Bool {
	return NewBool(m.Has(key))
}

func (m *Map) G_get(key Object) Object {
	return m.Get(key)
}

func (m *Map) G_set(key, value Object) {
	m.Set(key, value)
}

func (m *Map) G_remove(key Object) {
	m.Del(key)
}

func (m *Map) G_keys() *List {
	list := NewListSized(m.Len())
	index := 0

	for _, key := range m.Keys() {
		list.Set(index, key)
		index++
	}

	return list
}

func (m *Map) G_values() *List {
	list := NewListSized(m.Len())
	index := 0

	for _, value := range m.Values() {
		list.Set(index, value)
		index++
	}

	return list
}

func (m *Map) G_each(cb Object) {
	for hash, key := range m.Keys() {
		cb.G_invoke(NewVarArgs(key, m.values[hash]))
	}
}

func (m *Map) G_pop(key Object, fallback ...Object) Object {
	if !m.Has(key) {
		if len(fallback) == 0 {
			return NewNil()
		}

		return fallback[0]
	}

	defer m.Del(key)
	return m.Get(key)
}
