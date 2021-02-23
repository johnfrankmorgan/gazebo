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

func (m *BoundMethod) Call(args *Args) Object {
	return m.value.Call(args.ReflectValues())[0].Interface().(Object)
}
