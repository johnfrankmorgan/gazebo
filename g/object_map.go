package g

// ObjectMap is the underlying type of maps in gazebo
type ObjectMap struct {
	PartialObject
	value map[uint64]Object
	keys  map[uint64]Object
}

// NewObjectNil creates a new nil object
func NewObjectMap() *ObjectMap {
	return &ObjectMap{
		PartialObject: PartialObject{typ: TypeMap},
		value:         map[uint64]Object{},
		keys:          map[uint64]Object{},
	}
}

// Value satisfies the Object interface
func (m *ObjectMap) Value() interface{} {
	return m.value
}

// Call satisfies the Object interface
func (m *ObjectMap) Call(method string, args Args) Object {
	return m.call(m, method, args)
}

func (m *ObjectMap) Map() map[uint64]Object {
	return m.value
}

func (m *ObjectMap) Len() int {
	return len(m.value)
}

func (m *ObjectMap) hash(object Object) uint64 {
	return uint64(EnsureNumber(object.Call(Protocols.Hash, nil)).Float())
}

func (m *ObjectMap) Keys() *ObjectList {
	list := NewObjectList(make([]Object, len(m.keys)))
	idx := 0

	for _, key := range m.keys {
		list.Set(idx, key)
		idx++
	}

	return list
}

func (m *ObjectMap) Has(key Object) bool {
	_, ok := m.value[m.hash(key)]
	return ok
}

func (m *ObjectMap) Get(key Object) Object {
	if value, ok := m.value[m.hash(key)]; ok {
		return value
	}

	return NewObjectNil()
}

func (m *ObjectMap) Set(key Object, value Object) {
	hash := m.hash(key)

	m.value[hash] = value
	m.keys[hash] = value
}

func (m *ObjectMap) Delete(key Object) {
	hash := m.hash(key)
	delete(m.value, hash)
	delete(m.keys, hash)
}
