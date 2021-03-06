package expr

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
)

type Binary struct {
	Op    lexer.Token
	Left  Expression
	Right Expression
}

func (m *Binary) Compile() code.Code {
	fun, ok := protocols.BinaryOperators[m.Op.Value]
	if !ok {
		panic(fmt.Errorf("expr: unknown binary operator: %q", m.Op.Value))
	}

	code := m.Left.Compile()
	code = append(code, op.GetAttr.Ins(fun))
	code = append(code, m.Right.Compile()...)

	return append(code, op.CallFunc.Ins(1))
}
