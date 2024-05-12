package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprFalse(expr.False) {
	c.emit(opcode.LoadFalse)
}
