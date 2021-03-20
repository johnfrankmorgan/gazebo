package g

import (
	"encoding/gob"
	"hash/maphash"
	"strings"
)

var _ Object = &Map{}

type Map struct {
	Base
	hash   maphash.Hash
	keys   map[uint64]Object
	values map[uint64]Object
}

func NewMap() *Map {
	object := &Map{
		keys:   make(map[uint64]Object),
		values: make(map[uint64]Object),
	}

	object.SetType(TypeMap)
	object.SetSelf(object)
	object.SetAttrs(object)

	return object
}

func (m *Map) Value() interface{} {
	return m.values
}

func (m *Map) ToBool() *Bool {
	return NewBool(!m.IsEmpty())
}

func (m *Map) ToString() *String {
	var (
		buff strings.Builder
		pos  = 0
	)

	buff.WriteByte('{')

	for hash, key := range m.keys {
		pos++

		buff.WriteString(key.ToString().String())
		buff.WriteString(": ")
		buff.WriteString(m.values[hash].ToString().String())
		if pos < m.Len() {
			buff.WriteString(", ")
		}
	}

	buff.WriteByte('}')

	return NewString(buff.String())
}

func (m *Map) h(object Object) uint64 {
	defer m.hash.Reset()

	enc := gob.NewEncoder(&m.hash)

	if err := enc.Encode(object.Value()); err != nil {
		panic(err)
	}

	return m.hash.Sum64()
}

func (m *Map) Len() int {
	return len(m.keys)
}

func (m *Map) IsEmpty() bool {
	return m.Len() == 0
}

func (m *Map) Has(key Object) bool {
	_, ok := m.keys[m.h(key)]
	return ok
}

func (m *Map) Get(key Object) Object {
	if m.Has(key) {
		return m.values[m.h(key)]
	}

	return NewNil()
}

func (m *Map) Set(key, value Object) {
	hash := m.h(key)

	m.keys[hash] = key
	m.values[hash] = value
}

func (m *Map) Del(key Object) {
	hash := m.h(key)

	delete(m.keys, hash)
	delete(m.values, hash)
}

func (m *Map) Pop(key Object) Object {
	defer m.Del(key)

	return m.Get(key)
}
