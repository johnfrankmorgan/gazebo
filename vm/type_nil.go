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
