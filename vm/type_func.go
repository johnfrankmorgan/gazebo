package vm

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
)

var TypeFunc g.Type = &_func{}

type _func struct {
	g.Base
}

func (m *_func) Name() string {
	return "Func"
}

func (m *_func) Parent() g.Type {
	return g.TypeBase
}

func (m *_func) Methods() g.Methods {
	return g.Methods{
		protocols.Invoke: _func_invoke,
	}
}

func (m *_func) Value() interface{} {
	return m
}

func (m *_func) Type() g.Type {
	return g.TypeType
}

func _func_invoke(self g.Object, args *g.Args) g.Object {
	f := self.(*Func)

	if args.Len() < len(f.params) {
		panic(fmt.Errorf("expected %d arguments, got %d",
			len(f.params), args.Len()))
	}

	defer func(env *env) {
		f.vm.env = env
	}(f.vm.env)

	env := &env{parent: f.env}

	for i, param := range f.params {
		env.define(param, args.Get(i))
	}

	f.vm.env = env

	return f.vm.run(f.code)
}
