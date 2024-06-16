package vm

import (
	"fmt"
	"reflect"

	"github.com/johnfrankmorgan/gazebo/compile"
	"github.com/johnfrankmorgan/gazebo/compile/opcode"
	"github.com/johnfrankmorgan/gazebo/runtime"
)

type VM struct {
	Stack     Stack
	Variables *Variables
}

func New() *VM {
	variables := new(Variables)

	for _, builtin := range runtime.Builtins {
		variables.Set(string(builtin.Name()), builtin)
	}

	types := reflect.ValueOf(runtime.Types)
	for i := 0; i < types.NumField(); i++ {
		if field := types.Field(i); field.Type() == reflect.TypeOf((*runtime.Type)(nil)) {
			t := field.Interface().(*runtime.Type)

			variables.Set(string(t.Name), t)
		}
	}

	return &VM{
		Variables: variables,
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
			if runtime.Objects.Bool(vm.Stack.Pop()) {
				pc = arg - 1
			}

		case opcode.LoadLiteral:
			vm.Stack.Push(code.Literals[arg])

		case opcode.LoadName:
			value, ok := vm.Variables.Get(code.Idents[arg])
			if !ok {
				panic(runtime.Exc.NewUndefinedVariable(runtime.String(code.Idents[arg])))
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

			vm.Stack.Push(runtime.Objects.Attribute.Get(self, runtime.String(name)))

		case opcode.SetAttribute:
			self := vm.Stack.Pop()
			name := code.Idents[arg]
			value := vm.Stack.Pop()

			runtime.Objects.Attribute.Set(self, runtime.String(name), value)

		case opcode.Call:
			args := make(runtime.Tuple, arg)

			for i := range args {
				args[arg-i-1] = vm.Stack.Pop()
			}

			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Call(self, args))

		case opcode.GetIndex:
			index := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Index.Get(self, index))

		case opcode.SetIndex:
			index := vm.Stack.Pop()
			self := vm.Stack.Pop()
			value := vm.Stack.Pop()

			runtime.Objects.Index.Set(self, index, value)

		case opcode.LoadNil:
			vm.Stack.Push(runtime.Nil)

		case opcode.LoadFalse:
			vm.Stack.Push(runtime.False)

		case opcode.LoadTrue:
			vm.Stack.Push(runtime.True)

		// case opcode.Return:

		case opcode.UnaryNot:
			self := vm.Stack.Pop()
			vm.Stack.Push(!runtime.Objects.Bool(self))

		case opcode.UnaryPlus:
			self := vm.Stack.Pop()
			vm.Stack.Push(self)

		case opcode.UnaryMinus:
			self := vm.Stack.Pop()
			vm.Stack.Push(runtime.Objects.Unary.Negative(self))

		case opcode.BinaryAnd:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			result := runtime.Objects.Bool(self) && runtime.Objects.Bool(other)

			vm.Stack.Push(result)

		case opcode.BinaryOr:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			result := other

			if runtime.Objects.Bool(self) {
				result = self
			}

			vm.Stack.Push(result)

		case opcode.BinaryIs:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Is(self, other))

		case opcode.BinaryEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Equal(self, other))

		case opcode.BinaryNotEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.NotEqual(self, other))

		case opcode.BinaryLessThan:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Less(self, other))

		case opcode.BinaryLessThanOrEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.LessOrEqual(self, other))

		case opcode.BinaryGreaterThan:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Greater(self, other))

		case opcode.BinaryGreaterThanOrEqual:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.GreaterOrEqual(self, other))

		case opcode.BinaryIn:
			self := vm.Stack.Pop()
			other := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Contains(self, other))

		case opcode.BinaryAdd:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Add(self, other))

		case opcode.BinarySubtract:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Subtract(self, other))

		case opcode.BinaryMultiply:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Multiply(self, other))

		case opcode.BinaryDivide:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Divide(self, other))

		case opcode.BinaryModulo:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.Modulo(self, other))

		case opcode.BinaryBitwiseAnd:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.BitwiseAnd(self, other))

		case opcode.BinaryBitwiseOr:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.BitwiseOr(self, other))

		case opcode.BinaryBitwiseXor:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.BitwiseXor(self, other))

		case opcode.BinaryShiftLeft:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.ShiftLeft(self, other))

		case opcode.BinaryShiftRight:
			other := vm.Stack.Pop()
			self := vm.Stack.Pop()

			vm.Stack.Push(runtime.Objects.Binary.ShiftRight(self, other))

		default:
			panic(fmt.Errorf("runtime: unimplemented opcode: %v", op))
		}
	}

	if vm.Stack.Size() > 0 {
		return vm.Stack.Pop()
	}

	return runtime.Nil
}
