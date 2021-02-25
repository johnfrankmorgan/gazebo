package g

import "reflect"

type Args struct {
	Values []Object
}

func (m *Args) ReflectValues() []reflect.Value {
	values := make([]reflect.Value, len(m.Values))

	for i, value := range m.Values {
		values[i] = reflect.ValueOf(value)
	}

	return values
}
