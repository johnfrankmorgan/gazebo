package g

var TypeNil Type = &_nil{}

type _nil struct {
	Base
}

func (m *_nil) Name() string {
	return "Nil"
}

func (m *_nil) Parent() Type {
	return TypeBase
}

func (m *_nil) Methods() Methods {
	return Methods{}
}

func (m *_nil) Value() interface{} {
	return nil
}

func (m *_nil) Type() Type {
	return TypeType
}
