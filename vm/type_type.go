package vm

var _ Type = &TypeType{}

type TypeType struct {
	TypeBase
}

func (m *TypeType) Name() String {
	return NewString("Type")
}
