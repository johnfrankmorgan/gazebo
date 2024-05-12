package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtWhile(node stmt.While) {
	condition := c.label()
	c.compileExpr(node.Condition)

	body := label{}
	c.emit(opcode.JumpIfTrue, &body)

	end := label{}
	c.emit(opcode.Jump, &end)

	body.set(c.pc())
	c.compileStmt(node.Body)

	c.emit(opcode.Jump, condition.pc)
	end.set(c.pc())

	c.patch(c.pc(), condition.pc)
}
