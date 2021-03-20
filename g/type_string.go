package g

import "github.com/johnfrankmorgan/gazebo/g/protocols"

var TypeString Type = &_string{}

type _string struct {
	Base
}

func (m *_string) Name() string {
	return "String"
}

func (m *_string) Parent() Type {
	return TypeBase
}

func (m *_string) Methods() Methods {
	return Methods{
		protocols.Len: _string_len,
	}
}

func (m *_string) Value() interface{} {
	return m
}

func (m *_string) Type() Type {
	return TypeType
}

func _string_len(self Object, _ *Args) Object {
	return NewNumberFromInt(self.(*String).Len())
}
