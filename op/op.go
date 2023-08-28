package op

import "fmt"

//go:generate stringer -type=Opcode
type Opcode int

func Argument[T ~int](op T) Opcode {
	return -Opcode(op)
}

func (op Opcode) IsArgument() bool {
	return op <= 0
}

func (op Opcode) Value() int {
	return -int(op)
}

func (op Opcode) GoString() string {
	if op.IsArgument() {
		return fmt.Sprintf("op.Argument(%d)", op.Value())
	}

	return fmt.Sprintf("op.Opcode(%s(%d))", op, op)
}

const (
	_ Opcode = iota

	ExecuteChild

	LoadConstant
	LoadNull
	LoadFalse
	LoadTrue
	LoadName
	StoreName

	MakeFunc

	Jump
	RelativeJump
	RelativeJumpIfTrue
	RelativeJumpIfFalse

	UnaryNegate
	UnaryInvert

	BinaryAnd
	BinaryOr
	BinaryEqual
	BinaryNotEqual
	BinaryLess
	BinaryLessOrEqual
	BinaryGreater
	BinaryGreaterOrEqual
	BinaryAdd
	BinarySubtract
	BinaryMultiply
	BinaryDivide
	BinaryModulus

	Call
	Return
)

// keep in sync with the ast package's BinaryOp constants
var Binaries = [...]Opcode{
	0,
	BinaryAnd,
	BinaryOr,
	BinaryEqual,
	BinaryNotEqual,
	BinaryLess,
	BinaryLessOrEqual,
	BinaryGreater,
	BinaryGreaterOrEqual,
	BinaryAdd,
	BinarySubtract,
	BinaryMultiply,
	BinaryDivide,
	BinaryModulus,
}

// keep in sync with the ast package's UnaryOp constants
var Unaries = [...]Opcode{
	0,
	UnaryNegate,
	UnaryInvert,
}
