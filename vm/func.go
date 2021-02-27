package vm

import (
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Func{}

type Func struct {
	g.Base
	vm     *VM
	env    *env
	params []string
	code   compiler.Code
}

func NewFunc(vm *VM, env *env, params []string, code compiler.Code) *Func {
	object := &Func{
		vm:     vm,
		env:    env,
		params: params,
		code:   code,
	}
	object.SetSelf(object)
	return object
}

func (m *Func) Value() interface{} {
	return m.code
}

// GAZEBO FUNC OBJECT PROTOCOLS

func (m *Func) G_invoke(args *g.Args) g.Object {
	errors.ErrRuntime.Expect(
		args.Len() >= len(m.params),
		"expected %d arguments, got %d",
		len(m.params),
		args.Len(),
	)

	defer func(env *env) {
		m.vm.env = env
	}(m.vm.env)

	fenv := &env{parent: m.env}

	for i, param := range m.params {
		fenv.define(param, args.Get(i))
	}

	m.vm.env = fenv

	return m.vm.run(m.code)
}
