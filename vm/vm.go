package vm

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
)

type VM struct {
	env   *Env
	stack *Stack
	pc    int
}

func New() *VM {
	return &VM{
		env:   NewEnv(nil, nil),
		stack: NewStack(),
	}
}

func (m *VM) Run(code []compiler.Ins) Object {
	m.run(code)

	if m.stack.Size() > 0 {
		return m.stack.Peek()
	}

	return NewNil()
}

func (m *VM) run(code []compiler.Ins) {
	m.pc = 0

	for m.pc < len(code) {
		ins := code[m.pc]

		m.exec(ins)
		m.pc++
	}
}

func (m *VM) todo(ins compiler.Ins) {
	panic(
		fmt.Errorf(
			"opcode %q is not implemented yet: %v",
			ins.Op,
			ins,
		),
	)
}

func (m *VM) exec(ins compiler.Ins) {
	switch ins.Op {
	case op.LoadConst:
		m.stack.Push(NewObject(ins.Arg))

	case op.StoreName:
		m.env.Assign(ins.Arg.(string), m.stack.Peek())

	case op.LoadName:
		m.stack.Push(m.env.Lookup(ins.Arg.(string)))
	case op.Jump:
		m.pc = ins.Arg.(int)

	case op.RelJumpIfTrue:
		top := m.stack.Pop()
		if top.Type().ToBool(top, nil).Bool() {
			m.pc += ins.Arg.(int)
		}

	case op.RelJumpIfFalse:
		top := m.stack.Pop()
		if !top.Type().ToBool(top, nil).Bool() {
			m.pc += ins.Arg.(int)
		}

	case op.RelJump:
		m.pc += ins.Arg.(int)

	case op.BinEq:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Eq(self, Args{other}))

	case op.BinNEq:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().NEq(self, Args{other}))

	case op.BinAdd:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Add(self, Args{other}))

	case op.BinSub:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Sub(self, Args{other}))

	case op.BinMul:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Mul(self, Args{other}))

	case op.BinDiv:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Div(self, Args{other}))

	case op.BinLess:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Lt(self, Args{other}))

	case op.BinLessEq:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().LtE(self, Args{other}))

	case op.BinGreater:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().Gt(self, Args{other}))

	case op.BinGreaterEq:
		self := m.stack.Pop()
		other := m.stack.Pop()
		m.stack.Push(self.Type().GtE(self, Args{other}))

	case op.MakeFunction:
		def := ins.Arg.(*compiler.FuncDef)
		m.stack.Push(NewFunction(m, def.Args, def.Body))

	case op.Return:
		return

	case op.Call:
		argc := ins.Arg.(int)
		argv := make(Args, argc)

		for i := argc - 1; i >= 0; i-- {
			argv[i] = m.stack.Pop()
		}

		f := m.stack.Pop()
		m.stack.Push(f.Type().Call(f, argv))

	case op.AttrGet:
		top := m.stack.Pop()
		attr := NewString(ins.Arg.(string))
		m.stack.Push(top.Type().GetAttr(top, Args{attr}))

	default:
		m.todo(ins)
	}
}
