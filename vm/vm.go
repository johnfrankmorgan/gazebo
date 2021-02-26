package vm

import (
	"os"

	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules"
	"github.com/johnfrankmorgan/gazebo/g/modules/time"
)

// VM is the structure responsible for running code and keeping track of state
type VM struct {
	stack     *stack
	env       *env
	modules   map[string]modules.Module
	errhandle bool
}

// New creates a new VM
func New(argv ...string) *VM {
	env := new(env)

	env.define("out", g.NewWriter(os.Stdout))
	env.define("nil", g.NewNil())
	env.define("true", g.NewBool(true))
	env.define("false", g.NewBool(false))

	vm := &VM{
		stack:     new(stack),
		env:       env,
		errhandle: true,
		modules:   make(map[string]modules.Module),
	}

	time := time.NewTimeModule()
	vm.modules[time.Name()] = time

	return vm
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

		case op.GetName:
			name := ins.Arg.(string)
			m.stack.push(m.env.lookup(name))

		case op.SetName:
			name := ins.Arg.(string)
			if m.env.defined(name) {
				m.env.assign(name, m.stack.pop())
			} else {
				m.env.define(name, m.stack.pop())
			}

		case op.DelName:
			name := ins.Arg.(string)
			m.env.remove(name)

		case op.RelJump:
			pc += ins.Arg.(int)

		case op.RelJumpIfTrue:
			if m.stack.pop().G_bool().Bool() {
				pc += ins.Arg.(int)
			}

		case op.RelJumpIfFalse:
			if m.stack.pop().G_not().Bool() {
				pc += ins.Arg.(int)
			}

		case op.CallFunc:
			argc := ins.Arg.(int)
			args := g.NewArgs(make([]g.Object, argc))

			for i := 0; i < argc; i++ {
				args.Set(argc-i-1, m.stack.pop())
			}

			fun := m.stack.pop()
			m.stack.push(fun.G_invoke(args))

		case op.GetAttr:
			name := ins.Arg.(string)
			object := m.stack.pop()
			m.stack.push(object.G_getattr(g.NewString(name)))

		case op.SetAttr:
			name := ins.Arg.(string)
			value := m.stack.pop()
			object := m.stack.pop()
			m.stack.push(object.G_setattr(g.NewString(name), value))

		case op.DelAttr:
			name := ins.Arg.(string)
			object := m.stack.pop()
			m.stack.push(object.G_delattr(g.NewString(name)))

		case op.MakeList:
			length := ins.Arg.(int)
			values := make([]g.Object, length)

			for i := 0; i < length; i++ {
				values[length-i-1] = m.stack.pop()
			}

			m.stack.push(g.NewList(values))

		case op.Return:
			break loop

		case op.LoadModule:
			name := ins.Arg.(string)
			mod, ok := m.modules[name]
			errors.ErrRuntime.Expect(ok, "undefined module %q", name)
			m.env.define(name, mod)

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
