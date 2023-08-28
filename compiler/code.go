package compiler

import (
	"gazebo/ast"
	"gazebo/op"
)

type Code struct {
	Parent    *Code
	Opcodes   []op.Opcode
	Names     []string
	Constants []ast.Expression
	Children  []*Code
}

func (c *Code) Emit(opcode op.Opcode, args ...int) {
	c.Opcodes = append(c.Opcodes, opcode)

	c.EmitArgument(args...)
}

func (c *Code) EmitArgument(args ...int) {
	for _, arg := range args {
		c.Opcodes = append(c.Opcodes, op.Argument(arg))
	}
}

func (c *Code) Name(name string) int {
	for i, n := range c.Names {
		if n == name {
			return i
		}
	}

	c.Names = append(c.Names, name)

	return len(c.Names) - 1
}

func (c *Code) Constant(expression ast.Expression) int {
	c.Constants = append(c.Constants, expression)

	return len(c.Constants) - 1
}

func (c *Code) Child(child *Code) int {
	child.Parent = c

	c.Children = append(c.Children, child)

	return len(c.Children) - 1
}
