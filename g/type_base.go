package g

var TypeBase Type = &_base{}

type _base struct {
	Base
}

func (m *_base) Name() string {
	return "Base"
}

func (m *_base) Parent() Type {
	return nil
}

func (m *_base) Methods() Methods {
	return Methods{}
}

func (m *_base) Value() interface{} {
	return nil
}

func (m *_base) Type() Type {
	return TypeType
}
