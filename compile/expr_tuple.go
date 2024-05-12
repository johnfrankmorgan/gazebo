package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprTuple(node expr.Tuple) {
	for _, item := range node.Items {
		c.compileExpr(item)
	}

	c.emit(opcode.MakeTuple, len(node.Items))
}
