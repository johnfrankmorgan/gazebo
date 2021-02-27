package vm

import "github.com/johnfrankmorgan/gazebo/g"

var _ g.Object = &InternalObject{}

type InternalObject struct {
	g.Base
	value interface{}
}

func NewInternalObject(value interface{}) *InternalObject {
	object := &InternalObject{value: value}
	object.SetSelf(object)
	return object
}

func (m *InternalObject) Value() interface{} {
	return m.value
}
