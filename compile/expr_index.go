package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprIndex(node expr.Index) {
	c.compileExpr(node.Inner)
	c.compileExpr(node.Key)
	c.emit(opcode.GetIndex)
}
