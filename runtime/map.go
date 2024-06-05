package runtime

type Map struct {
	ht ht
}

func NewMap() *Map {
	return new(Map)
}

func (m *Map) Type() *Type {
	return Types.Map
}

func (m *Map) Bool() Bool {
	return m.Len() != 0
}

func (m *Map) Len() Int {
	return Int(m.ht.size)
}

func (m *Map) Contains(key Object) Bool {
	_, ok := m.Get(key)
	return ok
}

func (m *Map) Get(key Object) (Object, Bool) {
	return m.ht.get(key)
}

func (m *Map) GetIndex(index Object) Object {
	if value, ok := m.Get(index); ok {
		return value
	}

	panic(Exc.NewKeyNotFound(index))
}

func (m *Map) Set(key, value Object) {
	m.ht.set(key, value)
}
