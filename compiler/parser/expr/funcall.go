package expr

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type FunCall struct {
	Function Expression
	Args     []Expression
}

func (m *FunCall) Compile() code.Code {
	code := m.Function.Compile()
	argc := 0

	for _, arg := range m.Args {
		code = append(code, arg.Compile()...)
		argc++
	}

	return append(code, op.CallFunc.Ins(argc))
}
