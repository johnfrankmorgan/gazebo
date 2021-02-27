package g

// Attributes is a mapping of attribute names to Object values
type Attributes struct {
	values map[string]Object
}

func (m *Attributes) init() {
	if m.values == nil {
		m.values = make(map[string]Object)
	}
}

// Has returns whether an attribute exists
func (m *Attributes) Has(name string) bool {
	_, ok := m.values[name]
	return ok
}

// Get returns an attribute
func (m *Attributes) Get(name string) Object {
	if m.Has(name) {
		return m.values[name]
	}

	return NewNil()
}

// Set sets an attribute's value
func (m *Attributes) Set(name string, value Object) {
	m.init()

	m.values[name] = value
}

// Delete deletes an attribute
func (m *Attributes) Delete(name string) {
	delete(m.values, name)
}
