package vm

import (
	"os"

	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

// VM is the structure responsible for running code and keeping track of state
type VM struct {
	stack     *stack
	env       *env
	errhandle bool
}

// New creates a new VM
func New(argv ...string) *VM {
	env := new(env)

	env.define("out", g.NewWriter(os.Stdout))
	env.define("nil", g.NewNil())
	env.define("true", g.NewBool(true))
	env.define("false", g.NewBool(false))

	return &VM{
		stack:     new(stack),
		env:       env,
		errhandle: true,
	}
}

func (m *VM) DisableErrorHandling() {
	m.errhandle = false
}

// Run runs the provided code
func (m *VM) Run(code compiler.Code) (value g.Object, err error) {
	if m.errhandle {
		defer errors.Handle(&err)
	}

	value = m.run(code)
	return
}

func (m *VM) run(code compiler.Code) g.Object {
	var pc int

loop:
	for pc < len(code) {
		ins := code[pc]
		pc++

		switch ins.Opcode {
		case op.LoadConst:
			m.stack.push(g.NewObject(ins.Arg))

		case op.StoreName:
			name := ins.Arg.(string)
			if m.env.defined(name) {
				m.env.assign(name, m.stack.pop())
			} else {
				m.env.define(name, m.stack.pop())
			}

		case op.LoadName:
			name := ins.Arg.(string)
			m.stack.push(m.env.lookup(name))

		case op.CallFunc:
			argc := ins.Arg.(int)
			args := &g.Args{
				Values: make([]g.Object, argc),
			}

			for i := 0; i < argc; i++ {
				args.Values[argc-i-1] = m.stack.pop()
			}

			fun := m.stack.pop()
			m.stack.push(fun.(g.Callable).Call(args))

		case op.GetAttr:
			name := ins.Arg.(string)
			object := m.stack.pop()
			m.stack.push(object.GetAttr(name))

		case op.Return:
			break loop

		case op.NoOp:
			//

		default:
			assert.Unreached("unknown instruction: 0x%02x (%s) %#v", int(ins.Opcode), ins.Opcode.Name(), ins)
		}
	}

	if m.stack.size() > 0 {
		return m.stack.pop()
	}

	return nil
}
