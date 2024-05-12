package compile

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprUnary(node expr.Unary) {
	c.compileExpr(node.Right)

	switch node.Op {
	case expr.UnaryNot:
		c.emit(opcode.UnaryNot)

	case expr.UnaryPlus:
		c.emit(opcode.UnaryPlus)

	case expr.UnaryMinus:
		c.emit(opcode.UnaryMinus)

	default:
		panic(fmt.Errorf("compile: unknown unary operator: %v", node.Op))
	}
}
