package g

var TypeNumber Type = &_number{}

type _number struct {
	Base
}

func (m *_number) Name() string {
	return "Number"
}

func (m *_number) Parent() Type {
	return TypeBase
}

func (m *_number) Methods() Methods {
	return Methods{}
}

func (m *_number) Value() interface{} {
	return nil
}

func (m *_number) Type() Type {
	return TypeType
}
