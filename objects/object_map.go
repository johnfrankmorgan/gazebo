package objects

var _ Object = &Map{}

type Map struct {
	keys map[uint64]Object
	vals map[uint64]Object
}

func NewMap() *Map {
	return &Map{
		keys: make(map[uint64]Object),
		vals: make(map[uint64]Object),
	}
}

/* Object methods */

func (m *Map) GoVal() interface{} {
	panic("Map.GoVal()")
}

func (m *Map) Hash() interface{} {
	panic("Map.Hash()")
}

func (m *Map) Type() Type {
	return TypeMap
}

func (m *Map) ToString() *String {
	panic("Map.ToString()")
}

/* Methods */

func (m *Map) Has(key Object) *Bool {
	return TypeMap.G_has(m, key)
}

func (m *Map) Get(key Object) Object {
	return TypeMap.G_has(m, key)
}

func (m *Map) Pull(key Object) Object {
	return TypeMap.G_pull(m, key)
}

func (m *Map) Put(key, value Object) {
	TypeMap.G_put(m, key, value)
}

func (m *Map) Del(key Object) {
	TypeMap.G_del(m, key)
}
