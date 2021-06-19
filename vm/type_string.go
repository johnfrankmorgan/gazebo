package vm

var _ Type = &TypeString{}

type TypeString struct {
	TypeBase
}

func (m *TypeString) Name() *String {
	return NewString("String")
}
