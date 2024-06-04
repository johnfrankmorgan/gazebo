package runtime

import "fmt"

type Exception struct {
	typ     *Type
	message string
}

func (e *Exception) Type() *Type {
	return e.typ
}

func (e *Exception) Error() string {
	return string(e.Message())
}

func (e *Exception) Message() String {
	return Stringf("%s: %s", e.Type().Name, e.message)
}

func (e *Exception) Repr() String {
	return Stringf("%s(%q)", e.Type().Name, e.message)
}

func (e *Exception) String() String {
	return e.Message()
}

type _exc struct{}

var Exc _exc

func (_exc) NewUnimplemented(op string, t *Type) *Exception {
	return &Exception{
		typ:     Types.Exc.Unimplemented,
		message: fmt.Sprintf("%q not implemented for operand type %q", op, t.Name),
	}
}

func (_exc) NewUnimplementedUnary(op UnaryProtocol, t *Type) *Exception {
	return &Exception{
		typ:     Types.Exc.Unimplemented,
		message: fmt.Sprintf("unary %q not implemented for operand type %q", op, t.Name),
	}
}

func (_exc) NewUnimplementedBinary(op BinaryProtocol, stype, otype *Type) *Exception {
	return &Exception{
		typ:     Types.Exc.Unimplemented,
		message: fmt.Sprintf("binary %q not implemented for operand types %q and %q", op, stype.Name, otype.Name),
	}
}

func (_exc) NewInvalidType(got *Type, expected ...*Type) *Exception {
	names := make([]string, len(expected))

	for i, t := range expected {
		names[i] = string(t.Name)
	}

	return &Exception{
		typ:     Types.Exc.InvalidType,
		message: fmt.Sprintf("invalid type: got %s, expected %s", got.Name, names),
	}
}

func (_exc) NewInvalidAttributeGet(name String, t *Type) *Exception {
	return &Exception{
		typ:     Types.Exc.InvalidAttribute,
		message: fmt.Sprintf("invalid attribute: can't get %q for type %s", name, t.Name),
	}
}

func (_exc) NewInvalidAttributeSet(name String, t *Type) *Exception {
	return &Exception{
		typ:     Types.Exc.InvalidAttribute,
		message: fmt.Sprintf("invalid attribute: can't set %q for type %s", name, t.Name),
	}
}

func (_exc) NewInvalidIndex(index Object) *Exception {
	return &Exception{
		typ:     Types.Exc.InvalidIndex,
		message: fmt.Sprintf("invalid index: %v", index),
	}
}

func (_exc) NewOutOfBounds(index Object, length Int) *Exception {
	return &Exception{
		typ:     Types.Exc.OutOfBounds,
		message: fmt.Sprintf("out of bounds: %v for length %d", index, length),
	}
}

func (_exc) NewKeyNotFound(key Object) *Exception {
	return &Exception{
		typ:     Types.Exc.KeyNotFound,
		message: fmt.Sprintf("key not found: %v", key),
	}
}

func (_exc) NewUndefinedVariable(name String) *Exception {
	return &Exception{
		typ:     Types.Exc.UndefinedVariable,
		message: fmt.Sprintf("undefined variable: %q", name),
	}
}
