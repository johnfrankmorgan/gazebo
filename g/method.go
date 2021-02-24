package g

import (
	"reflect"

	"github.com/johnfrankmorgan/gazebo/assert"
)

var _ Object = &BoundMethod{}

type BoundMethod struct {
	AttrsNoOp
	value reflect.Value
}

func NewBoundMethod(value reflect.Value) *BoundMethod {
	return &BoundMethod{value: value}
}

func (m *BoundMethod) Value() interface{} {
	assert.Unreached()
	return nil
}

func (m *BoundMethod) CallMethod(name string, args *Args) Object {
	assert.Unreached()
	return nil
}

func (m *BoundMethod) Call(args *Args) Object {
	ret := m.value.Call(args.ReflectValues())
	if len(ret) == 0 {
		return NewNil()
	}

	assert.Len(ret, 1)

	return ret[0].Interface().(Object)
}
