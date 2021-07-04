package vm

var _ Type = &TypeNil{}

type TypeNil struct {
	TypeBase
}

func (m *TypeNil) Name() *String {
	return NewString("Nil")
}

func (m *TypeNil) IsNil(self Object, _ Args) *Bool {
	return NewBool(true)
}

func (m *TypeNil) ToString(self Object, _ Args) *String {
	return NewString("nil")
}

func (m *TypeNil) ToNumber(self Object, _ Args) *Number {
	return NewNumber(0.0)
}

func (m *TypeNil) ToBool(self Object, _ Args) *Bool {
	return NewBool(false)
}
