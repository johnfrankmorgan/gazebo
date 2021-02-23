package vm

import (
	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/compiler/op"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

// VM is the structure responsible for running code and keeping track of state
type VM struct {
	stack   *stack
	env     *env
	modules map[string]*modules.Module
}

// New creates a new VM
func New(argv ...string) *VM {
	env := new(env)

	for name, builtin := range g.Builtins() {
		env.define(name, builtin)
	}

	gargv := g.NewObjectList(nil)

	for _, arg := range argv {
		gargv.Append(g.NewObjectString(arg))
	}

	env.define("argv", gargv)

	return &VM{
		stack:   new(stack),
		env:     env,
		modules: modules.All(),
	}
}

// Run runs the provided code
func (m *VM) Run(code compiler.Code) (value g.Object, err error) {
	defer errors.Handle(&err)

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

		case op.RemoveName:
			name := ins.Arg.(string)
			m.env.remove(name)

		case op.CallFunc:
			argc := ins.Arg.(int)
			args := make(g.Args, argc)

			for i := 0; i < argc; i++ {
				args[argc-i-1] = m.stack.pop()
			}

			fun := m.stack.pop()

			switch fun.Type() {
			case g.TypeInternalFunc:
				m.stack.push(g.Invoke(fun, args))

			case g.TypeFunc:
				fun := g.EnsureFunc(fun)

				errors.ErrRuntime.ExpectLen(
					fun.Params(),
					len(args),
					"expected %d args, got %d",
					len(fun.Params()),
					len(args),
				)

				vmenv := m.env
				env := &env{
					parent: fun.Env().(*env),
				}

				for i, param := range fun.Params() {
					env.define(param, args[i])
				}

				m.env = env
				m.stack.push(m.run(fun.Code()))
				m.env = vmenv

			default:
				errors.ErrRuntime.Panic(
					"unexpected type called as function: gtypes.%s",
					fun.Type().Name,
				)
			}

		case op.RelJump:
			pc += ins.Arg.(int)

		case op.RelJumpIfTrue:
			condition := m.stack.pop()
			if g.IsTruthy(condition) {
				pc += ins.Arg.(int)
			}

		case op.RelJumpIfFalse:
			condition := m.stack.pop()
			if !g.IsTruthy(condition) {
				pc += ins.Arg.(int)
			}

		case op.PushValue:
			m.stack.push(g.NewObjectInternal(ins.Arg))

		case op.MakeFunc:
			body := m.stack.pop().Value().(compiler.Code)
			params := m.stack.pop().Value().([]string)
			m.stack.push(g.NewObjectFunc(params, body, m.env))

		case op.LoadModule:
			name := ins.Arg.(string)
			module, ok := m.modules[name]

			errors.ErrRuntime.Expect(ok, "undefined module: %s", name)

			m.env.define(name, module.Load())

		case op.MakeList:
			length := ins.Arg.(int)
			values := make([]g.Object, length)

			for i := 0; i < length; i++ {
				values[length-i-1] = m.stack.pop()
			}

			m.stack.push(g.NewObjectList(values))

		case op.IndexGet:
			index := m.stack.pop()
			value := m.stack.pop()
			m.stack.push(value.Call(protocols.Index, g.Args{index}))

		case op.AttributeGet:
			value := m.stack.pop()
			attr := g.NewObjectString(ins.Arg.(string))
			m.stack.push(value.Call(protocols.GetAttr, g.Args{attr}))

		case op.AttributeSet:
			value := m.stack.pop()
			object := m.stack.pop()
			attr := g.NewObjectString(ins.Arg.(string))
			m.stack.push(object.Call(protocols.SetAttr, g.Args{attr, value}))

		case op.NoOp:
			//

		case op.Return:
			break loop

		default:
			assert.Unreached("unknown instruction: 0x%02x (%s) %#v", int(ins.Opcode), ins.Opcode.Name(), ins)
		}
	}

	if m.stack.size() > 0 {
		return m.stack.pop()
	}

	return g.NewObjectNil()
}
