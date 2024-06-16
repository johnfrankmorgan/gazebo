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

func (i Int) Equal(other Object) Object {
	switch other := other.(type) {
	case Int:
		return Bool(i == other)

	case Float:
		return i.Float().Equal(other)
	}

	return Unimplemented
}

func (i Int) Less(other Object) Object {
	switch other := other.(type) {
	case Int:
		return Bool(i < other)

	case Float:
		return i.Float().Less(other)
	}

	return Unimplemented
}

func (i Int) Greater(other Object) Object {
	switch other := other.(type) {
	case Int:
		return Bool(i > other)

	case Float:
		return i.Float().Greater(other)
	}

	return Unimplemented
}

func (i Int) Add(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i + other

	case Float:
		return i.Float().Add(other)
	}

	return Unimplemented
}

func (i Int) Subtract(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i - other

	case Float:
		return i.Float().Subtract(other)
	}

	return Unimplemented
}

func (i Int) Multiply(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i * other

	case Float:
		return i.Float().Multiply(other)
	}

	return Unimplemented
}

func (i Int) Divide(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i.Float().Divide(other.Float())

	case Float:
		return i.Float().Divide(other)
	}

	return Unimplemented
}

func (i Int) Modulo(other Object) Object {
	switch other := other.(type) {
	case Int:
		return i % other

	case Float:
		return i.Float().Modulo(other)
	}

	return Unimplemented
}

func (i Int) BitwiseAnd(other Object) Object {
	if other, ok := other.(Int); ok {
		return i & other
	}

	return Unimplemented
}

func (i Int) BitwiseOr(other Object) Object {
	if other, ok := other.(Int); ok {
		return i | other
	}

	return Unimplemented
}

func (i Int) BitwiseXor(other Object) Object {
	if other, ok := other.(Int); ok {
		return i ^ other
	}

	return Unimplemented
}

func (i Int) ShiftLeft(other Object) Object {
	if other, ok := other.(Int); ok {
		return i << other
	}

	return Unimplemented
}

func (i Int) ShiftRight(other Object) Object {
	if other, ok := other.(Int); ok {
		return i >> other
	}

	return Unimplemented
}
