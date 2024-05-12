package runtime

type Map struct {
	ht ht
}

func NewMap() *Map {
	return new(Map)
}

func (m *Map) Get(key Object) (Object, Bool) {
	return m.ht.get(key)
}

func (m *Map) Set(key, value Object) {
	m.ht.set(key, value)
}
