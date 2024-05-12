package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtAssign(node stmt.Assign) {
	c.compileExpr(node.Expression)
	c.emit(opcode.StoreName, c.ident(node.Identifier))
}
