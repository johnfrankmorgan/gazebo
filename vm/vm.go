package vm

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/compile"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
	"github.com/johnfrankmorgan/gazebo/runtime"
)

type VM struct {
	Stack     Stack
	Variables *Variables
}

func New() *VM {
	return &VM{
		Variables: new(Variables),
	}
}

func (vm *VM) Exec(module *compile.Module) runtime.Object {
	code := module.Code

	for pc := 0; pc < len(code.Ops); pc++ {
		op := code.Ops[pc]
		arg := 0

		if op.HasArgument() {
			pc++
			arg = code.Ops[pc].Argument()
		}

		switch op {
		case opcode.Jump:
			pc = arg - 1

		case opcode.JumpIfTrue:
			if runtime.Truthy(vm.Stack.Pop()) {
				pc = arg - 1
			}

		case opcode.LoadLiteral:
			vm.Stack.Push(code.Literals[arg])

		case opcode.LoadName:
			value, ok := vm.Variables.Get(code.Idents[arg])
			if !ok {
				panic(fmt.Errorf("runtime: undefined variable: %v", code.Idents[arg]))
			}

			vm.Stack.Push(value)

		case opcode.StoreName:
			vm.Variables.Set(code.Idents[arg], vm.Stack.Pop())

		case opcode.MakeList:
			result := runtime.NewListWithLength(arg)

			for i := range result.Len() {
				result.Set(result.Len()-i-1, vm.Stack.Pop())
			}

			vm.Stack.Push(result)

		case opcode.MakeMap:
			result := runtime.NewMap()

			for range arg {
				value := vm.Stack.Pop()
				key := vm.Stack.Pop()

				result.Set(key, value)
			}

			vm.Stack.Push(result)

		case opcode.MakeTuple:
			result := make(runtime.Tuple, arg)

			for i := range result {
				result[arg-i-1] = vm.Stack.Pop()
			}

			vm.Stack.Push(result)

		case opcode.GetAttribute:
			self := vm.Stack.Pop()
			name := code.Idents[arg]

			vm.Stack.Push(runtime.GetAttribute(self, runtime.String(name)))

		case opcode.SetAttribute:
			self := vm.Stack.Pop()
			name := code.Idents[arg]
			value := vm.Stack.Pop()

			runtime.SetAttribute(self, runtime.String(name), value)

		case opcode.GetIndex:
			index := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.GetIndex(self, index))

		case opcode.SetIndex:
			index := vm.Stack.Pop()
			self := vm.Stack.Pop()
			value := vm.Stack.Pop()

			runtime.SetIndex(self, index, value)

		case opcode.LoadNil:
			vm.Stack.Push(runtime.Nil)

		case opcode.LoadFalse:
			vm.Stack.Push(runtime.False)

		case opcode.LoadTrue:
			vm.Stack.Push(runtime.True)

		// case opcode.Return:

		case opcode.UnaryNot:
			self := vm.Stack.Pop()
			vm.Stack.Push(!runtime.Truthy(self))

		case opcode.UnaryPlus:
			self := vm.Stack.Pop()
			vm.Stack.Push(self)

		case opcode.UnaryMinus:
			self := vm.Stack.Pop()
			vm.Stack.Push(runtime.Negative(self))

		case opcode.BinaryAnd:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			result := runtime.Truthy(self) && runtime.Truthy(other)

			vm.Stack.Push(result)

		case opcode.BinaryOr:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			result := other

			if runtime.Truthy(self) {
				result = self
			}

			vm.Stack.Push(result)

		case opcode.BinaryIs:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Is(self, other))

		case opcode.BinaryEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Equal(self, other))

		case opcode.BinaryNotEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.NotEqual(self, other))

		case opcode.BinaryLessThan:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Less(self, other))

		case opcode.BinaryLessThanOrEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.LessOrEqual(self, other))

		case opcode.BinaryGreaterThan:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Greater(self, other))

		case opcode.BinaryGreaterThanOrEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.GreaterOrEqual(self, other))

		case opcode.BinaryIn:
			self := vm.Stack.Pop()
			other := vm.Stack.Pop()

			vm.Stack.Push(runtime.Contains(self, other))

		case opcode.BinaryAdd:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Add(self, other))

		case opcode.BinarySubtract:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Subtract(self, other))

		case opcode.BinaryMultiply:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Multiply(self, other))

		case opcode.BinaryDivide:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Divide(self, other))

		case opcode.BinaryModulo:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Modulo(self, other))

		case opcode.BinaryBitwiseAnd:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.BitwiseAnd(self, other))

		case opcode.BinaryBitwiseOr:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.BitwiseAnd(self, other))

		case opcode.BinaryBitwiseXor:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.BitwiseAnd(self, other))

		case opcode.BinaryShiftLeft:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.LeftShift(self, other))

		case opcode.BinaryShiftRight:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.RightShift(self, other))

		default:
			panic(fmt.Errorf("runtime: unimplemented opcode: %v", op))
		}
	}

	if vm.Stack.Size() > 0 {
		return vm.Stack.Pop()
	}

	return runtime.Nil
}
