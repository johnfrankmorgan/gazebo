package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
)

type If struct {
	Condition Expression
	TruePath  Statement
	FalsePath Statement
}

func (m *If) Compile() code.Code {
	if m.FalsePath == nil {
		m.FalsePath = &Block{}
	}

	truecode := m.TruePath.Compile()
	falsecode := append(m.FalsePath.Compile(), op.RelJump.Ins(len(truecode)))
	condition := append(m.Condition.Compile(), op.RelJumpIfTrue.Ins(len(falsecode)))

	code := append(condition, falsecode...)
	return append(code, truecode...)
}
