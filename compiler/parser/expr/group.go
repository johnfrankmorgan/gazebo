package expr

import "github.com/johnfrankmorgan/gazebo/compiler/code"

type Group struct {
	Expr Expression
}

func (m *Group) Compile() code.Code {
	return m.Expr.Compile()
}
