package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprList(node expr.List) {
	for _, item := range node.Items {
		c.compileExpr(item)
	}

	c.emit(opcode.MakeList, len(node.Items))
}
