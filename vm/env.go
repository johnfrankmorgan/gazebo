package vm

type Env struct {
	parent *Env
	values map[string]Object
}

func NewEnv(values map[string]Object, parent *Env) *Env {
	if values == nil {
		values = map[string]Object{}
	}

	return &Env{values: values, parent: parent}
}

func (m *Env) Resolve(name string) *Env {
	if _, ok := m.values[name]; ok {
		return m
	}

	if m.parent != nil {
		return m.parent.Resolve(name)
	}

	return nil
}

func (m *Env) Defined(name string) bool {
	return m.Resolve(name) != nil
}

func (m *Env) Assign(name string, value Object) {
	if env := m.Resolve(name); env != nil {
		env.values[name] = value
		return
	}

	m.values[name] = value
}

func (m *Env) Lookup(name string) Object {
	if env := m.Resolve(name); env != nil {
		return env.values[name]
	}

	return NewNil()
}
