package runtime

type Int int64

func (i Int) Type() *Type {
	return Types.Int
}

func (i Int) Hash() uint64 {
	return hash(i)
}

func (i Int) Bool() Bool {
	return i != 0
}

func (i Int) Repr() String {
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

	panic(Exc.NewUnimplementedBinary(BinaryProtocolEqual, i.Type(), other.Type()))

}

func (i Int) Less(other Object) Bool {
	switch other := other.(type) {
	case Int:
		return i < other

	case Float:
		return i.Float().Less(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolLess, i.Type(), other.Type()))
}

func (i Int) Greater(other Object) Bool {
	switch other := other.(type) {
	case Int:
		return i > other

	case Float:
		return i.Float().Greater(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolGreater, i.Type(), other.Type()))
}

func (i Int) Add(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i + other

	case Float:
		return i.Float().Add(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolAdd, i.Type(), other.Type()))
}

func (i Int) Subtract(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i - other

	case Float:
		return i.Float().Subtract(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolSubtract, i.Type(), other.Type()))
}

func (i Int) Multiply(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i * other

	case Float:
		return i.Float().Multiply(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolMultiply, i.Type(), other.Type()))
}

func (i Int) Divide(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i.Float().Divide(other.Float())

	case Float:
		return i.Float().Divide(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolDivide, i.Type(), other.Type()))
}

func (i Int) Modulo(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i % other

	case Float:
		return i.Float().Modulo(other)
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolModulo, i.Type(), other.Type()))
}

func (i Int) BitwiseAnd(other Object) Object {
	if other, ok := other.(Int); ok {
		return i & other
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseAnd, i.Type(), other.Type()))
}

func (i Int) BitwiseOr(other Object) Object {
	if other, ok := other.(Int); ok {
		return i | other
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseOr, i.Type(), other.Type()))
}

func (i Int) BitwiseXor(other Object) Object {
	if other, ok := other.(Int); ok {
		return i ^ other
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolBitwiseXor, i.Type(), other.Type()))
}

func (i Int) ShiftLeft(other Object) Object {
	if other, ok := other.(Int); ok {
		return i << other
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolShiftLeft, i.Type(), other.Type()))
}

func (i Int) ShiftRight(other Object) Object {
	if other, ok := other.(Int); ok {
		return i >> other
	}

	panic(Exc.NewUnimplementedBinary(BinaryProtocolShiftRight, i.Type(), other.Type()))
}
