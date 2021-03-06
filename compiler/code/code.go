package code

import (
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/debug"
)

type Compiler interface {
	Compile() Code
}

type Code []op.Instruction

func (m Code) Dump() {
	for idx, ins := range m {
		debug.Printf(
			"%6d %18s (0x%02x) %v\n",
			idx,
			ins.Opcode.Name(),
			int(ins.Opcode),
			ins.Arg,
		)
	}
}
