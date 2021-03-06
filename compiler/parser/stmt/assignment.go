package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type Assignment struct {
	Name string
	Expr expr.Expression
}

func (m *Assignment) Compile() code.Code {
	return append(m.Expr.Compile(), op.SetName.Ins(m.Name))
}
