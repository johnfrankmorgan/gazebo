package vm

import "github.com/johnfrankmorgan/gazebo/compiler"

var _ Object = &Function{}

type Function struct {
	LazyAttributes
	vm   *VM
	args []string
	code []compiler.Ins
}

func NewFunction(vm *VM, args []string, code []compiler.Ins) *Function {
	return &Function{vm: vm, args: args, code: code}
}

func (m *Function) Value() interface{} {
	return m.code
}

func (m *Function) Type() Type {
	return Types.Function
}

func (m *Function) Call(args Args) Object {
	args.ExpectsExactly(len(m.args))

	defer func(env *Env, pc int) {
		m.vm.env = env
		m.vm.pc = pc
	}(m.vm.env, m.vm.pc)

	m.vm.env = NewEnv(nil, m.vm.env)

	for i, name := range m.args {
		m.vm.env.Assign(name, args[i])
	}

	m.vm.run(m.code)
	return m.vm.stack.Pop()
}
