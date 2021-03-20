package g

import "github.com/johnfrankmorgan/gazebo/g/protocols"

var TypeBoundMethod Type = &_bound_method{}

type _bound_method struct {
	Base
}

func (m *_bound_method) Name() string {
	return "BoundMethod"
}

func (m *_bound_method) Parent() Type {
	return TypeBase
}

func (m *_bound_method) Methods() Methods {
	return Methods{
		protocols.Invoke: _bound_method_invoke,
	}
}

func (m *_bound_method) Value() interface{} {
	return m
}

func (m *_bound_method) Type() Type {
	return TypeType
}

func _bound_method_invoke(self Object, args *Args) Object {
	return self.(*BoundMethod).Call(args)
}
