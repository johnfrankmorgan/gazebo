package vm

import (
	"fmt"
	"gazebo/ast"
	"gazebo/compiler"
	"gazebo/objects"
	"gazebo/op"
	"gazebo/util/ds"
	"gazebo/util/must"
	"strconv"
)

type VM struct {
	frames    *ds.Stack[*Frame]
	variables *Variables
}

func New() *VM {
	vm := &VM{
		frames:    ds.NewStack[*Frame](),
		variables: NewVariables(nil),
	}

	for name, builtin := range objects.Builtins {
		vm.variables.Store(name, builtin)
	}

	return vm
}

func (vm *VM) frame() *Frame {
	if vm.frames.Size() == 0 {
		return nil
	}

	return vm.frames.Peek()
}

func (vm *VM) Run(code *compiler.Code) *objects.Object {
	vm.frames.Push(NewFrame(vm.frame(), code))
	defer vm.frames.Pop()

	return vm.run(code)
}

func (vm *VM) run(code *compiler.Code) *objects.Object {
	for !vm.frame().ExecutionComplete() {
		opcode := vm.frame().NextOpcode()

		vm.exec(opcode)
	}

	if vm.frame().Stack.Size() > 0 {
		return vm.frame().Stack.Pop()
	}

	return objects.Singletons.Null.AsObject()
}

func (vm *VM) exec(opcode op.Opcode) {
	switch opcode {
	case op.ExecuteChild:
		child := vm.frame().Code.Children[vm.frame().NextArgument()]

		code := vm.frame().Code
		pc := vm.frame().PC

		vm.frame().Code = child
		vm.frame().PC = 0
		vm.variables = NewVariables(vm.variables)

		vm.run(vm.frame().Code)

		vm.frame().Code = code
		vm.frame().PC = pc
		vm.variables = vm.variables.Parent()

	case op.LoadConstant:
		constant := vm.frame().Code.Constants[vm.frame().NextArgument()]

		switch constant := constant.(type) {
		case *ast.Integer:
			value := must.Succeed(strconv.ParseInt(constant.Value, 0, 64))
			vm.frame().Stack.Push(objects.NewInteger(value).AsObject())

		case *ast.String:
			value := must.Succeed(strconv.Unquote(constant.Value))
			vm.frame().Stack.Push(objects.NewString(value).AsObject())

		default:
			panic(fmt.Errorf("vm: attempt to load unknown constant type (%T): %v", constant, constant))
		}

	case op.LoadNull:
		vm.frame().Stack.Push(objects.Singletons.Null.AsObject())

	case op.LoadFalse:
		vm.frame().Stack.Push(objects.Singletons.False.AsObject())

	case op.LoadTrue:
		vm.frame().Stack.Push(objects.Singletons.True.AsObject())

	case op.LoadName:
		name := vm.frame().Code.Names[vm.frame().NextArgument()]

		variables := vm.variables.Resolve(name)
		if variables == nil {
			panic(fmt.Errorf("vm: attempt to load undefined variable: %q", name))
		}

		vm.frame().Stack.Push(variables.Load(name))

	case op.StoreName:
		name := vm.frame().Code.Names[vm.frame().NextArgument()]

		variables := vm.variables.Resolve(name)
		if variables == nil {
			variables = vm.variables
		}

		variables.Store(name, vm.frame().Stack.Pop())

	case op.Jump:
		panic("todo")

	case op.RelativeJump:
		panic("todo")

	case op.RelativeJumpIfTrue:
		panic("todo")

	case op.RelativeJumpIfFalse:
		panic("todo")

	case op.UnaryNegate:
		panic("todo")

	case op.UnaryNot:
		panic("todo")

	case op.BinaryAnd:
		panic("todo")

	case op.BinaryOr:
		panic("todo")

	case op.BinaryEqual:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Equals(self, other).AsObject())

	case op.BinaryNotEqual:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		eq := self.Type.Equals(self, other).Value()

		vm.frame().Stack.Push(objects.Singletons.Bool(!eq).AsObject())

	case op.BinaryLess:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Less(self, other).AsObject())

	case op.BinaryLessOrEqual:
		panic("todo")

	case op.BinaryGreater:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Greater(self, other).AsObject())

	case op.BinaryGreaterOrEqual:
		panic("todo")

	case op.BinaryAdd:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Add(self, other))

	case op.BinarySubtract:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Subtract(self, other))

	case op.BinaryMultiply:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Multiply(self, other))

	case op.BinaryDivide:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Divide(self, other))

	case op.BinaryModulus:
		other := vm.frame().Stack.Pop()
		self := vm.frame().Stack.Pop()

		vm.frame().Stack.Push(self.Type.Modulus(self, other))

	case op.Call:
		self := vm.frame().Stack.Pop()
		args := make([]*objects.Object, 0, vm.frame().NextArgument())

		for len(args) != cap(args) {
			args = append(args, vm.frame().Stack.Pop())
		}

		vm.frame().Stack.Push(self.Type.Call(self, args...))

	default:
		panic(fmt.Errorf("vm: unimplemented opcode: %v", opcode))
	}
}
