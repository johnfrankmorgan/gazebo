package g

var TypeMap Type = &_map{}

type _map struct {
	Base
}

func (m *_map) Name() string {
	return "Map"
}

func (m *_map) Parent() Type {
	return TypeBase
}

func (m *_map) Methods() Methods {
	return Methods{}
}

func (m *_map) Value() interface{} {
	return m
}

func (m *_map) Type() Type {
	return TypeType
}
