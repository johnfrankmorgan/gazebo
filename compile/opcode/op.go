package opcode

import (
	"fmt"
	"strconv"
)

type Op int

func (op Op) GoString() string {
	if op.IsArgument() {
		return strconv.Itoa(op.Argument())
	}

	return "opcode." + op.String()
}

const argMask = 0x100000

func Argument[T ~int](op T) Op {
	return Op(op | argMask)
}

func (op Op) AsArgument() Op {
	return op | argMask
}

func (op Op) IsArgument() bool {
	return op&argMask != 0
}

func (op Op) Argument() int {
	if !op.IsArgument() {
		panic(fmt.Errorf("opcode: %v is not an argument", op))
	}

	return int(op &^ argMask)
}

func (op Op) HasArgument() bool {
	return op < _HasArgumentEnd
}

const (
	_ Op = iota

	Jump
	JumpIfTrue

	LoadLiteral
	LoadName
	StoreName

	MakeList
	MakeMap
	MakeTuple

	GetAttribute
	SetAttribute

	_HasArgumentEnd

	GetIndex
	SetIndex

	LoadNil
	LoadFalse
	LoadTrue

	Return

	UnaryNot
	UnaryPlus
	UnaryMinus

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
