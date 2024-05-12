package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprIdent(node expr.Ident) {
	c.emit(opcode.LoadName, c.ident(node.Name))
}
