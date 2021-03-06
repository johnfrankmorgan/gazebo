package expr

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/lexer"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
)

type Unary struct {
	Op    lexer.Token
	Right Expression
}

func (m *Unary) Compile() code.Code {
	fun, ok := protocols.UnaryOperators[m.Op.Value]
	if !ok {
		panic(fmt.Errorf("expr: unknown unary operator %q", m.Op.Value))
	}

	code := m.Right.Compile()
	code = append(code, op.GetAttr.Ins(fun))

	return append(code, op.CallFunc.Ins(0))
}
