package g

import "github.com/johnfrankmorgan/gazebo/g/protocols"

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
	return Methods{
		protocols.Add:      _number_add,
		protocols.Subtract: _number_sub,
		protocols.Multiply: _number_mul,
		protocols.Divide:   _number_div,
	}
}

func (m *_number) Value() interface{} {
	return m
}

func (m *_number) Type() Type {
	return TypeType
}

func _number_add(self Object, args *Args) Object {
	return NewNumber(self.(*Number).Float() + args.Get(0).ToNumber().Float())
}

func _number_sub(self Object, args *Args) Object {
	return NewNumber(self.(*Number).Float() - args.Get(0).ToNumber().Float())
}

func _number_mul(self Object, args *Args) Object {
	return NewNumber(self.(*Number).Float() * args.Get(0).ToNumber().Float())
}

func _number_div(self Object, args *Args) Object {
	return NewNumber(self.(*Number).Float() / args.Get(0).ToNumber().Float())
}
