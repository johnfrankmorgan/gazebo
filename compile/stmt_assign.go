package compile

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/ast/stmt"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileStmtAssign(node stmt.Assign) {
	c.compileExpr(node.Right)

	switch left := node.Left.(type) {
	case expr.Ident:
		c.emit(opcode.StoreName, c.ident(left.Name))

	case expr.Index:
		c.compileExpr(left.Inner)
		c.compileExpr(left.Key)
		c.emit(opcode.SetIndex)

	default:
		panic(fmt.Errorf("compile: unknown left assignment type: %T", node))
	}
}
