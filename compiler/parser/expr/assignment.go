package expr

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Assignment struct {
	Name string
	Expr Expression
}

func (m *Assignment) Compile() code.Code {
	return append(m.Expr.Compile(), op.SetName.Ins(m.Name))
}
