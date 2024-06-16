package runtime

import "math"

type Float float64

func (f Float) Type() *Type {
	return Types.Float
}

func (f Float) Hash() uint64 {
	return hash(f)
}

func (f Float) Bool() Bool {
	return f != 0
}

func (f Float) Repr() String {
	return Stringf("%f", f)
}

var epsilon = math.Nextafter(1, 2) - 1

func (f Float) Equal(other Object) Object {
	otherf := Float(0)

	switch other := other.(type) {
	case Float:
		otherf = other

	case Int:
		otherf = other.Float()

	default:
		return Unimplemented
	}

	return Bool(math.Abs(float64(f-otherf)) < epsilon)
}

func (f Float) Less(other Object) Object {
	switch other := other.(type) {
	case Float:
		return Bool(f < other)

	case Int:
		return Bool(f < other.Float())
	}

	return Unimplemented
}

func (f Float) Greater(other Object) Object {
	switch other := other.(type) {
	case Float:
		return Bool(f > other)

	case Int:
		return Bool(f > other.Float())
	}

	return Unimplemented
}

func (f Float) Add(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f + other

	case Int:
		return f + other.Float()
	}

	return Unimplemented
}

func (f Float) Subtract(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f - other

	case Int:
		return f - other.Float()
	}

	return Unimplemented
}

func (f Float) Multiply(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f * other

	case Int:
		return f * other.Float()
	}

	return Unimplemented
}

func (f Float) Divide(other Object) Object {
	switch other := other.(type) {
	case Float:
		return f / other

	case Int:
		return f / other.Float()
	}

	return Unimplemented
}

func (f Float) Modulo(other Object) Object {
	switch other := other.(type) {
	case Float:
		return Float(math.Mod(float64(f), float64(other)))

	case Int:
		return Float(math.Mod(float64(f), float64(other.Float())))
	}

	return Unimplemented
}
