package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtReturn(node stmt.Return) {
	if node.Expression != nil {
		c.compileExpr(node.Expression)
	} else {
		c.emit(opcode.LoadNil)
	}

	c.emit(opcode.Return)
}
