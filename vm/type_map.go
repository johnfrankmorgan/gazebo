package vm

var _ Type = &TypeMap{}

type TypeMap struct {
	TypeBase
}

func (m *TypeMap) Name() String {
	return NewString("Map")
}
