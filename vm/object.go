package vm

import "fmt"

type Object interface {
	Value() interface{}
	Type() Type
	Attrs() *Map
}

func NewObject(value interface{}) Object {
	switch value := value.(type) {
	case nil:
		return NewNil()

	case bool:
		return NewBool(value)

	case int:
		return NewNumber(float64(value))

	case int64:
		return NewNumber(float64(value))

	case float64:
		return NewNumber(value)

	case string:
		return NewString(value)
	}

	panic(
		fmt.Errorf(
			"couldn't convert type %T to Object: %q",
			value,
			value,
		),
	)
}
