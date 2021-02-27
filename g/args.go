package g

import "reflect"

var _ Object = &Args{}

type Args struct {
	List
}

func NewArgs(value []Object) *Args {
	object := &Args{List: *NewList(value)}
	object.SetSelf(object)
	return object
}

func NewVarArgs(values ...Object) *Args {
	return NewArgs(values)
}

func (m *Args) ReflectValues() []reflect.Value {
	values := make([]reflect.Value, m.Len())

	for i, value := range m.All() {
		values[i] = reflect.ValueOf(value)
	}

	return values
}
