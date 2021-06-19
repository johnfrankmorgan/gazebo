package vm

var _ Object = &Map{}

type Map struct {
	keys   map[uint64]Object
	values map[uint64]Object
}

func NewMap() *Map {
	return &Map{
		keys:   map[uint64]Object{},
		values: map[uint64]Object{},
	}
}

func (m *Map) Value() interface{} {
	return m.values
}

func (m *Map) Type() Type {
	return Types.Map
}

func (m *Map) Attrs() *Map {
	return m
}

func (m *Map) Has(key Object) bool {
	_, ok := m.values[key.Type().Hash(key, nil).UInt64()]
	return ok
}

func (m *Map) Get(key Object) Object {
	if value, ok := m.values[key.Type().Hash(key, nil).UInt64()]; ok {
		return value
	}

	return NewNil()
}

func (m *Map) Set(key, value Object) {
	hash := key.Type().Hash(key, nil).UInt64()
	m.keys[hash] = key
	m.values[hash] = value
}

func (m *Map) Del(key Object) {
	hash := key.Type().Hash(key, nil).UInt64()
	delete(m.keys, hash)
	delete(m.values, hash)
}
