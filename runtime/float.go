package runtime

import "math"

type Float float64

var FloatType = &Type{
	Name:   "Float",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Hash:   func(self Object) uint64 { return self.(Float).Hash() },
		Bool:   func(self Object) Bool { return self.(Float).Bool() },
		String: func(self Object) String { return self.(Float).String() },
	},
	Ops: TypeOps{
		Positive: func(self Object) Object { return self.(Float) },
		Negative: func(self Object) Object { return -self.(Float) },

		Equal:   func(self, other Object) Bool { return self.(Float).Equal(other) },
		Less:    func(self, other Object) Bool { return self.(Float).Less(other) },
		Greater: func(self, other Object) Bool { return self.(Float).Greater(other) },

		Add:      func(self, other Object) Object { return self.(Float).Add(other) },
		Subtract: func(self, other Object) Object { return self.(Float).Subtract(other) },
		Multiply: func(self, other Object) Object { return self.(Float).Multiply(other) },
		Divide:   func(self, other Object) Object { return self.(Float).Divide(other) },
	},
}

func (f Float) Type() *Type {
	return FloatType
}

func (f Float) Hash() uint64 {
	return hash(f)
}

func (f Float) Bool() Bool {
	return f != 0
}

func (f Float) String() String {
	return Stringf("%f", f)
}

var epsilon = math.Nextafter(1, 2) - 1

func (f Float) Equal(other Object) Bool {
	otherf := Float(0)

	switch other := other.(type) {
	case Float:
		otherf = other

	case Int:
		otherf = other.Float()

	default:
		panic(ErrUnimplemented)
	}

	return math.Abs(float64(f-otherf)) < epsilon
}

func (f Float) Less(other Object) Bool {
	switch other := other.(type) {
	case Float:
		return f < other

	case Int:
		return f < other.Float()
	}

	panic(ErrUnimplemented)
}

func (f Float) Greater(other Object) Bool {
	switch other := other.(type) {
	case Float:
		return f > other

	case Int:
		return f > other.Float()
	}

	panic(ErrUnimplemented)
}

func (f Float) Add(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f + other

	case Int:
		return f + other.Float()
	}

	panic(ErrUnimplemented)
}

func (f Float) Subtract(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f - other

	case Int:
		return f - other.Float()
	}

	panic(ErrUnimplemented)
}

func (f Float) Multiply(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f * other

	case Int:
		return f * other.Float()
	}

	panic(ErrUnimplemented)
}

func (f Float) Divide(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f / other

	case Int:
		return f / other.Float()
	}

	panic(ErrUnimplemented)
}

func (f Float) Modulo(other Object) Object {
	switch other := other.(type) {
	case Float:
		return Float(math.Mod(float64(f), float64(other)))

	case Int:
		return Float(math.Mod(float64(f), float64(other.Float())))
	}

	panic(ErrUnimplemented)
}
