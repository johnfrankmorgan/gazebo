package g

import "github.com/johnfrankmorgan/gazebo/assert"

type Object interface {
	Value() interface{}
	CallMethod(name string, args *Args) Object
	Attrs
	Protocols
}

func NewObject(value interface{}) Object {
	switch value := value.(type) {
	case nil:
		return NewNil()

	case bool:
		return NewBool(value)

	case int:
		return NewNumber(float64(value))

	case float64:
		return NewNumber(value)

	case string:
		return NewString(value)
	}

	assert.Unreached("type %T cannot be coerced into an object: %v", value, value)
	return nil
}

type Protocols interface {
	G_str() *String
	G_num() *Number
	G_bool() *Bool
	G_not() *Bool
	G_eq(Object) *Bool
	G_neq(Object) *Bool
	G_gt(Object) *Bool
	G_gte(Object) *Bool
	G_lt(Object) *Bool
	G_lte(Object) *Bool
}
