package vm

type LazyAttributes struct {
	attrs *Map
}

func (m *LazyAttributes) Attrs() *Map {
	if m.attrs == nil {
		m.attrs = NewMap()
	}

	return m.attrs
}
