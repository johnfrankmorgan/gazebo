package objects

import "reflect"

type Callable interface {
	Call(args Object) Object
	Object
}

type reflection_callable struct {
	f reflect.Value
}

func (m *reflection_callable) Call(args ...Object) Object {
	rargs := make([]reflect.Value, len(args))
	for i, arg := range args {
		rargs[i] = reflect.ValueOf(arg)
	}

	ret := m.f.Call(rargs)
	Ensure(len(ret) == 1)
	return ret[0].Interface().(Object)
}
