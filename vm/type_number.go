package vm

var _ Type = &TypeNumber{}

type TypeNumber struct {
	TypeBase
}

func (m *TypeNumber) Name() *String {
	return NewString("Number")
}
