package gazebo

import (
	"fmt"
	"log"

	"github.com/alecthomas/repr"
)

type vm struct {
	frames stack[vmframe]
}

type vmframe struct {
	parent *vmframe
	stack  stack[*Object]
	code   *code
	pc     int
	locals map[string]*Object
}

func (f *vmframe) global() *vmframe {
	for f.parent != nil {
		f = f.parent
	}

	return f
}

func (f *vmframe) decode() opcode {
	op := f.code.opcodes[f.pc]
	f.pc++
	return op
}

func (vm *vm) frame() *vmframe {
	return vm.frames.top()
}

func (vm *vm) run() {
	for vm.frame().pc < vm.frame().code.pc() {
		vm.exec(vm.frame().decode())
	}
}

func (vm *vm) exec(op opcode) {
	switch op {
	case opExecChild:
		child := vm.frame().decode()
		pc := vm.frame().pc

		vm.frame().code = vm.frame().code.children[child]
		vm.frame().pc = 0

		log.Println("executing child", vm.frame().code.opcodes)

		vm.run()

		vm.frame().code = vm.frame().code.parent
		vm.frame().pc = pc

	case opDump:
		obj := vm.frame().stack.pop()

		switch obj.Type {
		case Types.Type:
			obj := (*TypeObject)(obj.Ptr())
			repr.Println(obj, repr.IgnoreGoStringer())

		case Types.Bool:
			obj := (*BoolObject)(obj.Ptr())
			repr.Println(obj, repr.IgnoreGoStringer())

		case Types.Int:
			obj := (*IntObject)(obj.Ptr())
			repr.Println(obj, repr.IgnoreGoStringer())

		case Types.String:
			obj := (*StringObject)(obj.Ptr())
			repr.Println(obj, repr.IgnoreGoStringer())

		case Types.Null:
			obj := (*NullObject)(obj.Ptr())
			repr.Println(obj, repr.IgnoreGoStringer())

		case Types.Object:
			obj := (*Object)(obj.Ptr())
			repr.Println(obj, repr.IgnoreGoStringer())

		default:
			unreachable()
		}

	case opPrint:
		count := vm.frame().decode()
		for i := opcode(0); i < count; i++ {
			if i > 0 {
				fmt.Print(" ")
			}

			obj := vm.frame().stack.pop()
			fmt.Print(obj.Type.String(obj).Value())
		}

		fmt.Println()

	case opLoadConst:
		obj := vm.frame().code.constants[vm.frame().decode()]
		vm.frame().stack.push(obj)

	case opLoadNull:
		vm.frame().stack.push(Null.Object.AsObject())

	case opLoadTrue:
		vm.frame().stack.push(Bools.True.AsObject())

	case opLoadFalse:
		vm.frame().stack.push(Bools.False.AsObject())

	case opLoadName:
		name := vm.frame().code.names[vm.frame().decode()]

		obj := vm.frame().locals[name]
		if obj == nil {
			obj = vm.frame().global().locals[name]
		}

		if obj == nil {
			obj = Null.Object.AsObject()
		}

		vm.frame().stack.push(obj)

	case opStoreName:
		name := vm.frame().code.names[vm.frame().decode()]
		vm.frame().locals[name] = vm.frame().stack.pop()

	case opRelJump:
		vm.frame().pc += int(vm.frame().decode())

	case opRelJumpIfTrue:
		off := vm.frame().decode()

		if obj := vm.frame().stack.pop(); obj.Type.Bool(obj).Value() {
			vm.frame().pc += int(off)
		}

	case opRelJumpIfFalse:
		off := vm.frame().decode()

		if obj := vm.frame().stack.pop(); !obj.Type.Bool(obj).Value() {
			vm.frame().pc += int(off)
		}

	case opBinAdd:
		other := vm.frame().stack.pop()
		self := vm.frame().stack.pop()

		vm.frame().stack.push(self.Type.Add(self, other))

	case opBinSub:
		other := vm.frame().stack.pop()
		self := vm.frame().stack.pop()

		vm.frame().stack.push(self.Type.Sub(self, other))

	case opBinMul:
		other := vm.frame().stack.pop()
		self := vm.frame().stack.pop()

		vm.frame().stack.push(self.Type.Mul(self, other))

	case opBinDiv:
		other := vm.frame().stack.pop()
		self := vm.frame().stack.pop()

		vm.frame().stack.push(self.Type.Div(self, other))

	case opBinMod:
		other := vm.frame().stack.pop()
		self := vm.frame().stack.pop()

		vm.frame().stack.push(self.Type.Mod(self, other))

	default:
		assert(false, "unimplemented %s", op.GoString())
	}
}
