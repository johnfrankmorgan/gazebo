package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type SetAttr struct {
	Expr  expr.Expression
	Name  string
	Value expr.Expression
}

func (m *SetAttr) Compile() code.Code {
	code := m.Expr.Compile()
	code = append(code, m.Value.Compile()...)
	return append(code, op.SetAttr.Ins(m.Name))
}
