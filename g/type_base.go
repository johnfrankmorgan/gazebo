package g

import (
	"fmt"

	"github.com/johnfrankmorgan/gazebo/g/protocols"
)

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
	methods := Methods{
		protocols.Bool:   _base_bool,
		protocols.Not:    _base_not,
		protocols.And:    _base_and,
		protocols.Or:     _base_or,
		protocols.String: _base_string,
		protocols.Number: _base_number,
		"type":           _base_type,
		"is_a":           _base_is_a,
		"debug":          _base_debug,
	}

	for _, method := range protocols.All() {
		method := method

		if !methods.Has(method) {
			methods[method] = func(self Object, _ *Args) Object {
				panic(fmt.Errorf("%q not implemented for type %q", method, self.Type().Name()))
			}
		}
	}

	return methods
}

func (m *_base) Value() interface{} {
	return m
}

func (m *_base) Type() Type {
	return TypeType
}

func _base_bool(self Object, _ *Args) Object {
	return self.ToBool()
}

func _base_not(self Object, _ *Args) Object {
	return NewBool(!self.ToBool().Bool())
}

func _base_and(self Object, args *Args) Object {
	return NewBool(
		self.ToBool().Bool() && args.Get(0).ToBool().Bool(),
	)
}

func _base_or(self Object, args *Args) Object {
	if self.ToBool().Bool() {
		return self
	}

	return args.Get(0)
}

func _base_string(self Object, _ *Args) Object {
	return self.ToString()
}

func _base_number(self Object, _ *Args) Object {
	return self.ToNumber()
}

func _base_type(self Object, _ *Args) Object {
	return self.Type()
}

func _base_is_a(self Object, args *Args) Object {
	var is func(Type, Type) bool

	is = func(t1, t2 Type) bool {
		if t1 == t2 {
			return true
		}

		if t1.Parent() != nil {
			return is(t1.Parent(), t2)
		}

		return false
	}

	return NewBool(is(self.Type(), args.Get(0).Type()))
}

func _base_debug(self Object, _ *Args) Object {
	return NewStringf("%# v", self)
}
