package expr

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type SetAttr struct {
	Expr  Expression
	Name  string
	Value Expression
}

func (m *SetAttr) Compile() code.Code {
	code := m.Expr.Compile()
	code = append(code, m.Value.Compile()...)
	return append(code, op.SetAttr.Ins(m.Name))
}
