package g

import "github.com/johnfrankmorgan/gazebo/g/protocols"

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
	return Methods{
		protocols.Len: _list_len,
		"append":      _list_append,
		"prepend":     _list_prepend,
		"all":         _list_all,
		"any":         _list_any,
	}
}

func (m *_list) Value() interface{} {
	return m
}

func (m *_list) Type() Type {
	return TypeType
}

func _list_len(self Object, _ *Args) Object {
	return NewNumberFromInt(self.(*List).Len())
}

func _list_append(self Object, args *Args) Object {
	self.(*List).Append(args.Slice()...)
	return self
}

func _list_prepend(self Object, args *Args) Object {
	self.(*List).Prepend(args.Slice()...)
	return self
}

func _list_all(self Object, _ *Args) Object {
	for _, obj := range self.(*List).Slice() {
		if !obj.ToBool().Bool() {
			return NewBool(false)
		}
	}

	return NewBool(true)
}

func _list_any(self Object, _ *Args) Object {
	for _, obj := range self.(*List).Slice() {
		if obj.ToBool().Bool() {
			return NewBool(true)
		}
	}

	return NewBool(false)
}
