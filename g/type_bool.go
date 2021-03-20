package g

var TypeBool Type = &_bool{}

type _bool struct {
	Base
}

func (m *_bool) Name() string {
	return "Bool"
}

func (m *_bool) Parent() Type {
	return TypeBase
}

func (m *_bool) Methods() Methods {
	return Methods{}
}

func (m *_bool) Value() interface{} {
	return m
}

func (m *_bool) Type() Type {
	return TypeType
}
