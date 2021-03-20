package g

import (
	"fmt"
)

type Args struct {
	List
}

func NewArgs(args []Object) *Args {
	return &Args{List: *NewList(args)}
}

func NewArgsV(args ...Object) *Args {
	return NewArgs(args)
}

type Method func(Object, *Args) Object

type Methods map[string]Method

func (m Methods) Has(name string) bool {
	_, ok := m[name]
	return ok
}

func (m Methods) Get(name string) Method {
	return m[name]
}

func Resolve(t Type, name string) Method {
	if t.Methods().Has(name) {
		return t.Methods().Get(name)
	}

	if t.Parent() != nil {
		return Resolve(t.Parent(), name)
	}

	return nil
}

func Call(self Object, name string, args *Args) Object {
	method := Resolve(self.Type(), name)

	if method == nil {
		panic(fmt.Errorf("undefined method %q for type %q", name, self.Type().Name()))
	}

	return method(self, args)
}

func CallV(self Object, name string, args ...Object) Object {
	return Call(self, name, NewArgs(args))
}
