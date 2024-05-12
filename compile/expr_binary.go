package compile

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/ast/expr"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
)

func (c *compiler) compileExprBinary(node expr.Binary) {
	c.compileExpr(node.Left)
	c.compileExpr(node.Right)

	switch node.Op {
	case expr.BinaryAnd:
		c.emit(opcode.BinaryAnd)

	case expr.BinaryOr:
		c.emit(opcode.BinaryOr)

	case expr.BinaryIs:
		c.emit(opcode.BinaryIs)

	case expr.BinaryEqual:
		c.emit(opcode.BinaryEqual)

	case expr.BinaryNotEqual:
		c.emit(opcode.BinaryNotEqual)

	case expr.BinaryLessThan:
		c.emit(opcode.BinaryLessThan)

	case expr.BinaryLessThanOrEqual:
		c.emit(opcode.BinaryLessThanOrEqual)

	case expr.BinaryGreaterThan:
		c.emit(opcode.BinaryGreaterThan)

	case expr.BinaryGreaterThanOrEqual:
		c.emit(opcode.BinaryGreaterThanOrEqual)

	case expr.BinaryIn:
		c.emit(opcode.BinaryIn)

	case expr.BinaryAdd:
		c.emit(opcode.BinaryAdd)

	case expr.BinarySubtract:
		c.emit(opcode.BinarySubtract)

	case expr.BinaryMultiply:
		c.emit(opcode.BinaryMultiply)

	case expr.BinaryDivide:
		c.emit(opcode.BinaryDivide)

	case expr.BinaryModulo:
		c.emit(opcode.BinaryModulo)

	case expr.BinaryBitwiseAnd:
		c.emit(opcode.BinaryBitwiseAnd)

	case expr.BinaryBitwiseOr:
		c.emit(opcode.BinaryBitwiseOr)

	case expr.BinaryBitwiseXor:
		c.emit(opcode.BinaryBitwiseXor)

	case expr.BinaryShiftLeft:
		c.emit(opcode.BinaryShiftLeft)

	case expr.BinaryShiftRight:
		c.emit(opcode.BinaryShiftRight)

	default:
		panic(fmt.Errorf("compile: unknown binary operator: %v", node.Op))
	}
}
