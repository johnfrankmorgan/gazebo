package g

import (
	"reflect"

	"github.com/johnfrankmorgan/gazebo/assert"
)

var _ Object = &BoundMethod{}

type BoundMethod struct {
	Base
	value reflect.Value
}

func NewBoundMethod(value reflect.Value) *BoundMethod {
	object := &BoundMethod{value: value}
	object.SetSelf(object)
	return object
}

func (m *BoundMethod) Value() interface{} {
	return m.value
}

func (m *BoundMethod) G_invoke(args *Args) Object {
	ret := m.value.Call(args.ReflectValues())
	if len(ret) == 0 {
		return NewNil()
	}

	assert.Len(ret, 1, "too many return values from %v", m.value)

	return ret[0].Interface().(Object)
}
