package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprTrue(expr.True) {
	c.emit(opcode.LoadTrue)
}
