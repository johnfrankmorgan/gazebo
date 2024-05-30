package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprAttr(node expr.Attr) {
	c.compileExpr(node.Inner)
	c.emit(opcode.GetAttribute, c.ident(node.Name))
}
