package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprCall(node expr.Call) {
	c.compileExpr(node.Target)

	for _, arg := range node.Arguments {
		c.compileExpr(arg)
	}

	c.emit(opcode.Call, len(node.Arguments))
}
