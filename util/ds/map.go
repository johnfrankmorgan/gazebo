package ds

type Map[K comparable, V any] struct {
	values map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		values: make(map[K]V),
	}
}

func (m *Map[K, V]) Has(key K) bool {
	_, ok := m.values[key]
	return ok
}

func (m *Map[K, V]) Get(key K) V {
	return m.values[key]
}

func (m *Map[K, V]) Set(key K, value V) {
	m.values[key] = value
}
