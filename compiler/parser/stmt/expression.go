package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type Expression struct {
	Expr expr.Expression
}

func (m *Expression) Compile() code.Code {
	return m.Expr.Compile()
}
