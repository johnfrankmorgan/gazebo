package expr

import (
	"github.com/johnfrankmorgan/gazebo/ast"
)

type Binary struct {
	base

	Op    BinaryOp
	Left  ast.Expr
	Right ast.Expr
}

type BinaryOp int

func (op BinaryOp) GoString() string {
	return "expr." + op.String()
}

const (
	_ BinaryOp = iota

	BinaryAnd
	BinaryOr

	BinaryIs
	BinaryEqual
	BinaryNotEqual
	BinaryLessThan
	BinaryLessThanOrEqual
	BinaryGreaterThan
	BinaryGreaterThanOrEqual

	BinaryIn

	BinaryAdd
	BinarySubtract
	BinaryMultiply
	BinaryDivide
	BinaryModulo

	BinaryBitwiseAnd
	BinaryBitwiseOr
	BinaryBitwiseXor

	BinaryShiftLeft
	BinaryShiftRight
)
