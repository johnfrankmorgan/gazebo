package runtime

import (
	"errors"
	"fmt"
	"slices"
	"unsafe"
)

type Object interface {
	Type() *Type
}

var ObjectType = &Type{
	Name: "Object",

	Protocols: TypeProtocols{
		Bool:   func(Object) Bool { return True },
		String: func(self Object) String { return Stringf("%v", self) },
	},

	Ops: TypeOps{
		Equal: Is,

		GetAttribute: func(self Object, name String) Object {
			for t := self.Type(); t != nil; t = t.Parent {
				attr, ok := t.Attributes[name]
				if !ok {
					continue
				}

				if attr.Get != nil {
					return attr.Get(self)
				}
			}

			panic(fmt.Errorf("%w: can't get %q", ErrInvalidAttribute, name))
		},

		SetAttribute: func(self Object, name String, value Object) {
			for t := self.Type(); t != nil; t = t.Parent {
				attr, ok := t.Attributes[name]
				if !ok {
					continue
				}

				if attr.Set != nil {
					attr.Set(self, value)
					return
				}
			}

			panic(fmt.Errorf("%w: can't set %q", ErrInvalidAttribute, name))
		},
	},

	Attributes: TypeAttributes{
		"type": Attribute{
			Get: func(self Object) Object { return self.Type() },
		},
	},
}

func Is(a, b Object) Bool {
	return a == b
}

var (
	ErrUnimplemented    = errors.New("runtime: unimplemented")
	ErrInvalidAttribute = errors.New("runtime: invalid attribute")
)

func unimplemented(op, kind string, t *Type) error {
	return fmt.Errorf("%w: %s %s not implemented for type %s", ErrUnimplemented, op, kind, t.Name)
}

func Truthy(a Object) Bool {
	for t := a.Type(); t != nil; t = t.Parent {
		if t.Protocols.Bool != nil {
			return t.Protocols.Bool(a)
		}
	}

	panic(unimplemented("bool", "protocol", a.Type()))
}

func Hash(a Object) uint64 {
	for t := a.Type(); t != nil; t = t.Parent {
		if t.Protocols.Hash != nil {
			return t.Protocols.Hash(a)
		}
	}

	panic(unimplemented("hash", "protocol", a.Type()))
}

func unop[T Object](op string, off uintptr, a Object) T {
	for t := a.Type(); t != nil; t = t.Parent {
		op := unsafe.Pointer(uintptr(unsafe.Pointer(&t.Ops)) + off)

		if fn := *(*func(Object) T)(op); fn != nil {
			return fn(a)
		}
	}

	panic(unimplemented(op, "operation", a.Type()))
}

func Positive(a Object) Object {
	return unop[Object]("positive", unsafe.Offsetof(TypeOps{}.Positive), a)
}

func Negative(a Object) Object {
	return unop[Object]("negative", unsafe.Offsetof(TypeOps{}.Negative), a)
}

func GetAttribute(a Object, name String) Object {
	for t := a.Type(); t != nil; t = t.Parent {
		if t.Ops.GetAttribute != nil {
			return t.Ops.GetAttribute(a, name)
		}
	}

	panic(unimplemented("get attribute", "operation", a.Type()))
}

func SetAttribute(a Object, name String, value Object) {
	for t := a.Type(); t != nil; t = t.Parent {
		if t.Ops.SetAttribute != nil {
			t.Ops.SetAttribute(a, name, value)
			return
		}
	}

	panic(unimplemented("set attribute", "operation", a.Type()))
}

func _binop[T Object](op string, off uintptr, a, b Object) (result T, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			if rerr, ok := rerr.(error); ok && errors.Is(rerr, ErrUnimplemented) {
				err = rerr
				return
			}

			panic(err)
		}
	}()

	for t := a.Type(); t != nil; t = t.Parent {
		op := unsafe.Pointer(uintptr(unsafe.Pointer(&t.Ops)) + off)

		if fn := *(*func(Object, Object) T)(op); fn != nil {
			return fn(a, b), nil
		}
	}

	return result, unimplemented(op, "operation", a.Type())
}

func binop[T Object](op string, off uintptr, a, b Object) (result T) {
	isCommutative := func(off uintptr) bool {
		return slices.Contains([]uintptr{
			unsafe.Offsetof(TypeOps{}.Equal),
			unsafe.Offsetof(TypeOps{}.Add),
			unsafe.Offsetof(TypeOps{}.Multiply),
			unsafe.Offsetof(TypeOps{}.BitwiseAnd),
			unsafe.Offsetof(TypeOps{}.BitwiseOr),
			unsafe.Offsetof(TypeOps{}.BitwiseXor),
		}, off)
	}

	result, err := _binop[T](op, off, a, b)

	if err != nil {
		// FIXME: add support for noncommutative operations
		if errors.Is(err, ErrUnimplemented) && isCommutative(off) {
			if result, err = _binop[T](op, off, b, a); err != nil {
				panic(err)
			}
		}
	}

	return result
}

func Equal(a, b Object) Bool {
	if Is(a, b) {
		return True
	}

	return binop[Bool]("equal", unsafe.Offsetof(TypeOps{}.Equal), a, b)
}

func NotEqual(a, b Object) Bool {
	return !Equal(a, b)
}

func Less(a, b Object) Bool {
	return binop[Bool]("less", unsafe.Offsetof(TypeOps{}.Less), a, b)
}

func LessOrEqual(a, b Object) Bool {
	return Less(a, b) || Equal(a, b)
}

func Greater(a, b Object) Bool {
	return binop[Bool]("greater", unsafe.Offsetof(TypeOps{}.Greater), a, b)
}

func GreaterOrEqual(a, b Object) Bool {
	return Greater(a, b) || Equal(a, b)
}

func Contains(a, b Object) Bool {
	return binop[Bool]("contains", unsafe.Offsetof(TypeOps{}.Contains), a, b)
}

func Add(a, b Object) Object {
	return binop[Object]("add", unsafe.Offsetof(TypeOps{}.Add), a, b)
}

func Subtract(a, b Object) Object {
	return binop[Object]("subtract", unsafe.Offsetof(TypeOps{}.Subtract), a, b)
}

func Multiply(a, b Object) Object {
	return binop[Object]("multiply", unsafe.Offsetof(TypeOps{}.Multiply), a, b)
}

func Divide(a, b Object) Object {
	return binop[Object]("divide", unsafe.Offsetof(TypeOps{}.Divide), a, b)
}

func Modulo(a, b Object) Object {
	return binop[Object]("modulo", unsafe.Offsetof(TypeOps{}.Modulo), a, b)
}

func BitwiseAnd(a, b Object) Object {
	return binop[Object]("bitwise and", unsafe.Offsetof(TypeOps{}.BitwiseAnd), a, b)
}

func BitwiseOr(a, b Object) Object {
	return binop[Object]("bitwise or", unsafe.Offsetof(TypeOps{}.BitwiseOr), a, b)
}

func BitwiseXor(a, b Object) Object {
	return binop[Object]("bitwise xor", unsafe.Offsetof(TypeOps{}.BitwiseXor), a, b)
}

func LeftShift(a, b Object) Object {
	return binop[Object]("left shift", unsafe.Offsetof(TypeOps{}.LeftShift), a, b)
}

func RightShift(a, b Object) Object {
	return binop[Object]("right shift", unsafe.Offsetof(TypeOps{}.RightShift), a, b)
}

func GetIndex(object, index Object) Object {
	for t := object.Type(); t != nil; t = t.Parent {
		if t.Ops.GetIndex != nil {
			return t.Ops.GetIndex(object, index)
		}
	}

	panic(unimplemented("get index", "operation", object.Type()))
}

func SetIndex(object, index, value Object) {
	for t := object.Type(); t != nil; t = t.Parent {
		if t.Ops.SetIndex != nil {
			t.Ops.SetIndex(object, index, value)
			return
		}
	}

	panic(unimplemented("set index", "operation", object.Type()))

}
