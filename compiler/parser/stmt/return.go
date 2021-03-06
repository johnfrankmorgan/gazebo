package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type Return struct {
	Expr expr.Expression
}

func (m *Return) Compile() code.Code {
	return append(m.Expr.Compile(), op.Return.Ins(nil))
}
