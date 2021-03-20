package g

var TypeType Type = &_type{}

type _type struct {
	Base
}

func (m *_type) Name() string {
	return "Type"
}

func (m *_type) Parent() Type {
	return nil
}

func (m *_type) Methods() Methods {
	return Methods{}
}

func (m *_type) Value() interface{} {
	return m
}

func (m *_type) Type() Type {
	return TypeType
}
