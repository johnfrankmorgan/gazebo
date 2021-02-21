package g

import (
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

// EnsureNil asserts that an Object is an ObjectNil
func EnsureNil(value Object) *ObjectNil {
	errors.ErrRuntime.Expect(
		value.Type() == TypeNil,
		"expected type Nil got %s",
		value.Type().Name,
	)

	return value.(*ObjectNil)
}

// EnsureBool asserts that an Object is an ObjectBool
func EnsureBool(value Object) *ObjectBool {
	errors.ErrRuntime.Expect(
		value.Type() == TypeBool,
		"expected type Bool got %s",
		value.Type().Name,
	)

	return value.(*ObjectBool)
}

// EnsureNumber asserts that an Object is an ObjectNumber
func EnsureNumber(value Object) *ObjectNumber {
	errors.ErrRuntime.Expect(
		value.Type() == TypeNumber,
		"expected type Number got %s",
		value.Type().Name,
	)

	return value.(*ObjectNumber)
}

// EnsureString asserts that an Object is an ObjectString
func EnsureString(value Object) *ObjectString {
	errors.ErrRuntime.Expect(
		value.Type() == TypeString,
		"expected type String got %s",
		value.Type().Name,
	)

	return value.(*ObjectString)
}

// EnsureList asserts that an Object is an ObjectList
func EnsureList(value Object) *ObjectList {
	errors.ErrRuntime.Expect(
		value.Type() == TypeList,
		"expected type List got %s",
		value.Type().Name,
	)

	return value.(*ObjectList)
}

// EnsureInternalFunc asserts that an Object is an ObjectInternalFunc
func EnsureInternalFunc(value Object) *ObjectInternalFunc {
	errors.ErrRuntime.Expect(
		value.Type() == TypeInternalFunc,
		"expected type InternalFunc got %s",
		value.Type().Name,
	)

	return value.(*ObjectInternalFunc)
}

// EnsureFunc asserts that an Object is an ObjectFunc
func EnsureFunc(value Object) *ObjectFunc {
	errors.ErrRuntime.Expect(
		value.Type() == TypeFunc,
		"expected type Func got %s",
		value.Type().Name,
	)

	return value.(*ObjectFunc)
}

// EnsureInternal asserts that an Object is an ObjectInternal
func EnsureInternal(value Object) *ObjectInternal {
	errors.ErrRuntime.Expect(
		value.Type() == TypeInternal,
		"expected type Internal got %s",
		value.Type().Name,
	)

	return value.(*ObjectInternal)
}

// IsTruthy determines if the provided Object is truthy
func IsTruthy(object Object) bool {
	return EnsureBool(object.Call(protocols.Bool, nil)).Bool()
}

// ToString returns an Object's string value
func ToString(object Object) string {
	return EnsureString(object.Call(protocols.String, nil)).String()
}

// ToFloat returns an Object's float value
func ToFloat(object Object) float64 {
	return EnsureNumber(object.Call(protocols.Number, nil)).Float()
}

// ToInt returns an Object's int value
func ToInt(object Object) int {
	return EnsureNumber(object.Call(protocols.Number, nil)).Int()
}

// Invoke calls an Object
func Invoke(object Object, args Args) Object {
	return object.Call(protocols.Invoke, args)
}
