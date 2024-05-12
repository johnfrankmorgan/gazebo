package compile

import (
	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
	"github.com/johnfrankmorgan/gazebo/runtime"
)

func (c *compiler) compileExprFloat(node expr.Float) {
	c.emit(opcode.LoadLiteral, c.literal(runtime.Float(node.Value)))
}
