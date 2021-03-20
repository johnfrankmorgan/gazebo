package g

var TypeList Type = &_list{}

type _list struct {
	Base
}

func (m *_list) Name() string {
	return "List"
}

func (m *_list) Parent() Type {
	return TypeBase
}

func (m *_list) Methods() Methods {
	return Methods{}
}

func (m *_list) Value() interface{} {
	return nil
}

func (m *_list) Type() Type {
	return TypeType
}
