package expr

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
)

type GetAttr struct {
	Expr Expression
	Name string
}

func (m *GetAttr) Compile() code.Code {
	return append(m.Expr.Compile(), op.GetAttr.Ins(m.Name))
}
