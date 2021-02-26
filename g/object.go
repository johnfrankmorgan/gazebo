package g

import "github.com/johnfrankmorgan/gazebo/assert"

func NewObject(value interface{}) Object {
	switch value := value.(type) {
	case nil:
		return NewNil()

	case bool:
		return NewBool(value)

	case int:
		return NewNumberFromInt(value)

	case float64:
		return NewNumber(value)

	case string:
		return NewString(value)

	case []Object:
		return NewList(value)
	}

	assert.Unreached("type %T cannot be coerced into an object: %v", value, value)
	return nil
}

type Object interface {
	Value() interface{}
	CallMethod(name string, args *Args) Object
	Attrs
	Protocols
}

type Attrs interface {
	HasAttr(string) bool
	GetAttr(string) Object
	SetAttr(string, Object)
	DelAttr(string)
}

type Protocols interface {
	G_repr() *String
	G_str() *String
	G_num() *Number
	G_bool() *Bool
	G_not() *Bool
	G_len() *Number
	G_inverse() Object
	G_and(Object) *Bool
	G_or(Object) Object
	G_contains(Object) *Bool
	G_add(Object) Object
	G_sub(Object) Object
	G_mul(Object) Object
	G_div(Object) Object
	G_eq(Object) *Bool
	G_neq(Object) *Bool
	G_gt(Object) *Bool
	G_gte(Object) *Bool
	G_lt(Object) *Bool
	G_lte(Object) *Bool
	G_hasattr(*String) *Bool
	G_getattr(*String) Object
	G_setattr(*String, Object) Object
	G_delattr(*String) Object
	G_invoke(*Args) Object
}
