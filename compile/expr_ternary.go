package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprTernary(node expr.Ternary) {
	c.compileExpr(node.Condition)

	consequence := label{}
	c.emit(opcode.JumpIfTrue, &consequence)

	c.compileExpr(node.Alternative)

	end := label{}
	c.emit(opcode.Jump, &end)

	consequence.set(c.pc())
	c.compileExpr(node.Consequence)

	end.set(c.pc())
}
