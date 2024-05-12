package compile

import (
	"fmt"
	"slices"

	"github.com/johnfrankmorgan/gazebo/compile/opcode"
	"github.com/johnfrankmorgan/gazebo/grammar"
	"github.com/johnfrankmorgan/gazebo/runtime"
)

type Module struct {
	Program *grammar.Program
	Code    *Code
}

type Code struct {
	Ops      []opcode.Op
	Idents   []string
	Literals []runtime.Object
}

func Compile(program *grammar.Program) *Module {
	module := &Module{
		Program: program,
		Code:    &Code{},
	}

	compiler := compiler{
		code: module.Code,
	}

	for _, stmt := range program.Statements {
		compiler.compileStmt(stmt)
	}

	return module
}

type compiler struct {
	code *Code
}

type label struct {
	c  *compiler
	pc int
}

func (l label) set(value int) {
	l.c.code.Ops[l.pc] = opcode.Argument(value)
}

func (c *compiler) label() label {
	return label{
		c:  c,
		pc: c.pc(),
	}
}

func (c *compiler) pc() int {
	return len(c.code.Ops)
}

func (c *compiler) emit(op opcode.Op, args ...any) {
	c.code.Ops = append(c.code.Ops, op)

	for _, arg := range args {
		switch arg := arg.(type) {
		case int:
			c.code.Ops = append(c.code.Ops, opcode.Argument(arg))

		case opcode.Op:
			c.code.Ops = append(c.code.Ops, arg)

		case *label:
			*arg = c.label()
			c.code.Ops = append(c.code.Ops, opcode.Argument(pendingLabel))

		default:
			panic(fmt.Errorf("compile: unknown argument type: %T", arg))
		}
	}
}

const (
	pendingLabel    opcode.Op = 0x1fffffff
	pendingBreak    opcode.Op = 0x1ffffffe
	pendingContinue opcode.Op = 0x1ffffffd
)

func (c *compiler) patch(brk, cont int) {
	for i := 0; i < len(c.code.Ops); i++ {
		if c.code.Ops[i] != opcode.Jump {
			continue
		}

		i++

		switch op := c.code.Ops[i]; op {
		case pendingBreak:
			c.code.Ops[i] = opcode.Argument(brk)

		case pendingContinue:
			c.code.Ops[i] = opcode.Argument(cont)
		}
	}
}

func (c *compiler) ident(ident string) int {
	index := slices.Index(c.code.Idents, ident)

	if index < 0 {
		index = len(c.code.Idents)
		c.code.Idents = append(c.code.Idents, ident)
	}

	return index
}

func (c *compiler) literal(literal runtime.Object) int {
	c.code.Literals = append(c.code.Literals, literal)

	return len(c.code.Literals) - 1
}
