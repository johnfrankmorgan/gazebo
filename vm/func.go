package vm

import (
	"fmt"
	"strings"

	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/g"
)

type Func struct {
	g.Base
	vm     *VM
	env    *env
	params []string
	code   code.Code
}

func NewFunc(vm *VM, env *env, params []string, code code.Code) *Func {
	object := &Func{
		vm:     vm,
		env:    env,
		params: params,
		code:   code,
	}

	object.SetType(TypeFunc)
	object.SetSelf(object)

	return object
}

func (m *Func) Value() interface{} {
	return m
}

func (m *Func) ToString() *g.String {
	var buff strings.Builder

	buff.WriteString(m.Type().Name())
	buff.WriteString(fmt.Sprintf("@%p", m))
	buff.WriteByte('(')

	for i, param := range m.params {
		buff.WriteString(param)

		if i < len(m.params)-1 {
			buff.WriteString(", ")
		}
	}

	buff.WriteByte(')')

	return g.NewString(buff.String())
}
