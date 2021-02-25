package g

import (
	"reflect"

	"github.com/johnfrankmorgan/gazebo/assert"
)

var _ Object = &BoundMethod{}

type BoundMethod struct {
	Base
	value reflect.Value
	self  Object
	name  string
}

func NewBoundMethod(value reflect.Value) *BoundMethod {
	object := &BoundMethod{value: value}
	object.self = object
	return object
}

func (m *BoundMethod) Value() interface{} {
	assert.Unreached()
	return nil
}

func (m *BoundMethod) Call(args *Args) Object {
	if !m.value.IsValid() {
		m.value = m.Method(m.self, m.name).value
	}

	ret := m.value.Call(args.ReflectValues())
	if len(ret) == 0 {
		return NewNil()
	}

	assert.Len(ret, 1)

	return ret[0].Interface().(Object)
}
