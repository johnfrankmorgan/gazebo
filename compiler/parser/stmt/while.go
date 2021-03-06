package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/compiler/parser/expr"
)

type While struct {
	Loop
	Condition expr.Expression
	Body      Statement
}

func (m *While) Compile() code.Code {
	body := m.Body.Compile()
	cond := append(m.Condition.Compile(), op.RelJumpIfFalse.Ins(len(body)+1))
	body = append(body, op.RelJump.Ins(-len(body)-len(cond)-1))
	return m.FillLoop(append(cond, body...))
}
