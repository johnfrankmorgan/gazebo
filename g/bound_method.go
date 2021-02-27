package g

import (
	"fmt"
	"reflect"

	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/errors"
)

var _ Object = &BoundMethod{}

type BoundMethod struct {
	Base
	typ   string
	name  string
	value reflect.Value
}

func NewBoundMethod(typ string, name string, value reflect.Value) *BoundMethod {
	object := &BoundMethod{
		typ:   typ,
		name:  name,
		value: value,
	}
	object.SetSelf(object)
	return object
}

func (m *BoundMethod) Value() interface{} {
	return m.value
}

func (m *BoundMethod) Name() string {
	return fmt.Sprintf("%s.%s", m.typ, m.name)
}

// GAZEBO BOUND METHOD OBJECT PROTOCOLS

func (m *BoundMethod) G_repr() *String {
	return NewStringf("<%T>(%s.%s)", m, m.typ, m.name)
}

func (m *BoundMethod) G_invoke(args *Args) Object {
	if !m.value.Type().IsVariadic() {
		errors.ErrRuntime.Expect(
			args.Len() == m.value.Type().NumIn(),
			"%s expects %d arguments, got %d",
			m.Name(),
			m.value.Type().NumIn(),
			args.Len(),
		)
	}

	ret := m.value.Call(args.ReflectValues())
	if len(ret) == 0 || ret[0].Interface() == nil {
		return NewNil()
	}

	assert.Len(ret, 1, "too many return values from %v", m.value)

	return ret[0].Interface().(Object)
}

// GAZEBO BOUND METHOD OBJECT METHODS

func (m *BoundMethod) G_type() *String {
	return NewString(m.typ)
}

func (m *BoundMethod) G_name() *String {
	return NewString(m.name)
}
