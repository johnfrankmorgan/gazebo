package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtIf(node stmt.If) {
	c.compileExpr(node.Condition)

	consequence := label{}
	c.emit(opcode.JumpIfTrue, &consequence)

	if node.Alternative != nil {
		c.compileStmt(node.Alternative)
	}

	end := label{}
	c.emit(opcode.Jump, &end)

	consequence.set(c.pc())
	c.compileStmt(node.Consequence)

	end.set(c.pc())
}
