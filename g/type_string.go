package g

var TypeString Type = &_string{}

type _string struct {
	Base
}

func (m *_string) Name() string {
	return "String"
}

func (m *_string) Parent() Type {
	return nil
}

func (m *_string) Methods() Methods {
	return Methods{}
}

func (m *_string) Value() interface{} {
	return nil
}

func (m *_string) Type() Type {
	return TypeType
}
