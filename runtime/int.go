package runtime

type Int int64

var IntType = &Type{
	Name:   "Int",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Hash:   func(self Object) uint64 { return self.(Int).Hash() },
		Bool:   func(self Object) Bool { return self.(Int).Bool() },
		String: func(self Object) String { return self.(Int).String() },
	},
	Ops: TypeOps{
		Positive: func(self Object) Object { return self.(Int) },
		Negative: func(self Object) Object { return -self.(Int) },

		Equal:   func(self, other Object) Bool { return self.(Int).Equal(other) },
		Less:    func(self, other Object) Bool { return self.(Int).Less(other) },
		Greater: func(self, other Object) Bool { return self.(Int).Greater(other) },

		Add:      func(self, other Object) Object { return self.(Int).Add(other) },
		Subtract: func(self, other Object) Object { return self.(Int).Subtract(other) },
		Multiply: func(self, other Object) Object { return self.(Int).Multiply(other) },
		Divide:   func(self, other Object) Object { return self.(Int).Divide(other) },
		Modulo:   func(self, other Object) Object { return self.(Int).Modulo(other) },

		BitwiseAnd: func(self, other Object) Object { return self.(Int).BitwiseAnd(other) },
		BitwiseOr:  func(self, other Object) Object { return self.(Int).BitwiseOr(other) },
		BitwiseXor: func(self, other Object) Object { return self.(Int).BitwiseXor(other) },

		LeftShift:  func(self, other Object) Object { return self.(Int).LeftShift(other) },
		RightShift: func(self, other Object) Object { return self.(Int).RightShift(other) },
	},
}

func (i Int) Type() *Type {
	return IntType
}

func (i Int) Hash() uint64 {
	return hash(i)
}

func (i Int) Bool() Bool {
	return i != 0
}

func (i Int) String() String {
	return Stringf("%d", i)
}

func (i Int) Float() Float {
	return Float(i)
}

func (i Int) Equal(other Object) Bool {
	switch other := other.(type) {
	case Int:
		return i == other

	case Float:
		return i.Float().Equal(other)
	}

	panic(ErrUnimplemented)

}

func (i Int) Less(other Object) Bool {
	switch other := other.(type) {
	case Int:
		return i < other

	case Float:
		return i.Float().Less(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) Greater(other Object) Bool {
	switch other := other.(type) {
	case Int:
		return i > other

	case Float:
		return i.Float().Greater(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) Add(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i + other

	case Float:
		return i.Float().Add(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) Subtract(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i - other

	case Float:
		return i.Float().Subtract(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) Multiply(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i * other

	case Float:
		return i.Float().Multiply(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) Divide(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i / other

	case Float:
		return i.Float().Divide(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) Modulo(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i % other

	case Float:
		return i.Float().Modulo(other)
	}

	panic(ErrUnimplemented)
}

func (i Int) BitwiseAnd(other Object) Object {
	if other, ok := other.(Int); ok {
		return i & other
	}

	panic(ErrUnimplemented)
}

func (i Int) BitwiseOr(other Object) Object {
	if other, ok := other.(Int); ok {
		return i | other
	}

	panic(ErrUnimplemented)
}

func (i Int) BitwiseXor(other Object) Object {
	if other, ok := other.(Int); ok {
		return i ^ other
	}

	panic(ErrUnimplemented)
}

func (i Int) LeftShift(other Object) Object {
	if other, ok := other.(Int); ok {
		return i << other
	}

	panic(ErrUnimplemented)
}

func (i Int) RightShift(other Object) Object {
	if other, ok := other.(Int); ok {
		return i >> other
	}

	panic(ErrUnimplemented)
}
