package vm

import (
	"gazebo/compiler"
	"gazebo/objects"
	"gazebo/op"
	"gazebo/util/ds"
)

type Frame struct {
	Parent *Frame
	Stack  *ds.Stack[*objects.Object]
	Code   *compiler.Code
	PC     int
	Result *objects.Object
}

func NewFrame(parent *Frame, code *compiler.Code) *Frame {
	return &Frame{
		Parent: parent,
		Stack:  ds.NewStack[*objects.Object](),
		Code:   code,
	}
}

func (f *Frame) NextOpcode() op.Opcode {
	opcode := f.Code.Opcodes[f.PC]
	f.PC++
	return opcode
}

func (f *Frame) NextArgument() int {
	return f.NextOpcode().Value()
}

func (f *Frame) ExecutionComplete() bool {
	return f.PC >= len(f.Code.Opcodes)
}
