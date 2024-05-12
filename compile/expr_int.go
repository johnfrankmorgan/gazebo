package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
	"github.com/johnfrankmorgan/gazebo/runtime"
)

func (c *compiler) compileExprInt(node expr.Int) {
	c.emit(opcode.LoadLiteral, c.literal(runtime.Int(node.Value)))
}
