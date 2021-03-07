package vm

import (
	"io/ioutil"
	"path/filepath"

	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/compiler"
	"github.com/johnfrankmorgan/gazebo/compiler/code"
	"github.com/johnfrankmorgan/gazebo/compiler/code/op"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/modules"
	"github.com/johnfrankmorgan/gazebo/g/modules/os"
	"github.com/johnfrankmorgan/gazebo/g/modules/testing"
)

type VM struct {
	stack   *stack
	env     *env
	modules map[string]modules.Module
}

func New() *VM {
	vm := &VM{
		stack:   new(stack),
		env:     new(env),
		modules: make(map[string]modules.Module),
	}

	for _, mod := range modules.All() {
		vm.modules[mod.Name()] = mod
	}

	stdout := vm.modules["os"].(*os.OSModule).Stdout
	stderr := vm.modules["os"].(*os.OSModule).Stderr

	vm.modules["testing"].(*testing.TestingModule).SetOutput(
		stdout,
		stderr,
	)

	vm.env.define("nil", g.NewNil())
	vm.env.define("true", g.NewBool(true))
	vm.env.define("false", g.NewBool(false))
	vm.env.define("out", stdout)
	return vm
}

func (m *VM) GetModule(name string) modules.Module {
	return m.modules[name]
}

func (m *VM) RunFile(path string) (g.Object, error) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	code, err := compiler.Compile(string(source))
	if err != nil {
		return nil, err
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	m.env.define("__file", g.NewString(path))

	return m.Run(code)
}

func (m *VM) Run(code code.Code) (value g.Object, err error) {
	defer errors.Handle(&err)

	value = m.run(code)
	return
}

func (m *VM) run(instructions code.Code) g.Object {
	var pc int

loop:
	for pc < len(instructions) {
		ins := instructions[pc]

		if debug.Enabled() {
			debug.Printf("INSTRUCTION ")
			instructions[pc : pc+1].Dump()
			debug.Printf("\n")
			m.stack.dump()
		}

		pc++

		switch ins.Opcode {
		case op.PushValue:
			m.stack.push(NewInternalObject(ins.Arg))

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

		case op.MakeFunc:
			code := m.stack.pop().Value().(code.Code)
			params := m.stack.pop().Value().([]string)
			m.stack.push(NewFunc(m, m.env, params, code))

		case op.MakeList:
			length := ins.Arg.(int)
			list := g.NewListSized(length)

			for i := 0; i < length; i++ {
				list.Set(length-i-1, m.stack.pop())
			}

			m.stack.push(list)

		case op.MakeMap:
			length := ins.Arg.(int)
			hmap := g.NewMap()

			for i := 0; i < length; i++ {
				value := m.stack.pop()
				key := m.stack.pop()

				hmap.Set(key, value)
			}

			m.stack.push(hmap)

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
