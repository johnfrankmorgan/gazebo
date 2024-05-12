package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
	"github.com/johnfrankmorgan/gazebo/runtime"
)

func (c *compiler) compileExprString(node expr.String) {
	c.emit(opcode.LoadLiteral, c.literal(runtime.String(node.Value)))
}
