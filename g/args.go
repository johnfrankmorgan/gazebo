package g

import "reflect"

var _ Object = &Args{}

type Args struct {
	List
}

func NewArgs(value []Object) *Args {
	return &Args{List: *NewList(value)}
}

func (m *Args) ReflectValues() []reflect.Value {
	values := make([]reflect.Value, m.Len())

	for i, value := range m.All() {
		values[i] = reflect.ValueOf(value)
	}

	return values
}
