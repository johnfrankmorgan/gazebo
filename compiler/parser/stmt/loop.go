package stmt

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
)

type Loop struct{}

func (_ Loop) FillBreak(code code.Code) code.Code {
	for i, ins := range code {
		if ins.Opcode != op.Placeholder || ins.Arg.(int) != op.PlaceholderBreak {
			continue
		}

		code[i] = op.RelJump.Ins(len(code) - i - 1)
	}

	return code
}

func (_ Loop) FillContinue(code code.Code) code.Code {
	for i, ins := range code {
		if ins.Opcode != op.Placeholder || ins.Arg.(int) != op.PlaceholderContinue {
			continue
		}

		code[i] = op.RelJump.Ins(-i - 1)
	}

	return code
}

func (m Loop) FillLoop(code code.Code) code.Code {
	return m.FillContinue(m.FillBreak(code))
}
