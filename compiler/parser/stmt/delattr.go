package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type DelAttr struct {
	Expr expr.Expression
	Name string
}

func (m *DelAttr) Compile() code.Code {
	return append(m.Expr.Compile(), op.DelAttr.Ins(m.Name))
}
