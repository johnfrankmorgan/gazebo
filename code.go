package gazebo

import "fmt"

type code struct {
	parent    *code
	opcodes   []opcode
	names     []string
	constants []*Object
	children  []*code
}

//go:generate stringer -type=opcode
type opcode int

const (
	opInvalid opcode = iota

	opPrint
	opDump
	opExecChild

	opLoadConst
	opLoadNull
	opLoadTrue
	opLoadFalse
	opLoadName
	opStoreName

	opJump
	opRelJump
	opRelJumpIfTrue
	opRelJumpIfFalse

	opReturn
	opReturnNull

	opUnNegate
	opUnNot

	opBinAnd
	opBinOr
	opBinEqual
	opBinNotEqual
	opBinLess
	opBinLessEqual
	opBinGreater
	opBinGreaterEqual
	opBinAdd
	opBinSub
	opBinMul
	opBinDiv
	opBinMod
)

func (op opcode) GoString() string {
	return fmt.Sprintf("%s: %d", op, op)
}

func (c *code) pc() int {
	return len(c.opcodes)
}

func (c *code) emit(op opcode, args ...int) {
	c.opcodes = append(c.opcodes, op)

	for _, arg := range args {
		c.opcodes = append(c.opcodes, opcode(arg))
	}
}

func (c *code) name(name string) int {
	for i, n := range c.names {
		if n == name {
			return i
		}
	}

	c.names = append(c.names, name)
	return len(c.names) - 1
}

func (c *code) constant(obj *Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

func (c *code) child(child *code) int {
	child.parent = c
	c.children = append(c.children, child)
	return len(c.children) - 1
}
