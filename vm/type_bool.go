package vm

var _ Type = &TypeBool{}

type TypeBool struct {
	TypeBase
}

func (m *TypeBool) Name() *String {
	return NewString("Bool")
}
