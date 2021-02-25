package g

import (
	"reflect"

	"github.com/johnfrankmorgan/gazebo/errors"
)

type ObjectHelper struct {
	Attrs map[string]Object
}

func (m *ObjectHelper) Method(object Object, name string) *BoundMethod {
	method := reflect.ValueOf(object).MethodByName("G_" + name)

	if method.IsValid() {
		return &BoundMethod{value: method}
	}

	return nil
}

func (m *ObjectHelper) CallMethod(object Object, name string, args *Args) Object {
	if method := m.Method(object, name); method != nil {
		return method.Call(args)
	}

	errors.ErrRuntime.Panic("undefined method: %s", name)
	return nil
}

func (m *ObjectHelper) HasAttr(object Object, name string) bool {
	errors.ErrRuntime.Panic("not implemented: HasAttr")
	return false
}

func (m *ObjectHelper) GetAttr(object Object, name string) Object {
	if method := m.Method(object, name); method != nil {
		return method
	}

	errors.ErrRuntime.Panic("not implemented: GetAttr")
	return nil
}

func (m *ObjectHelper) SetAttr(object Object, name string, value Object) {
	errors.ErrRuntime.Panic("not implemented: SetAttr")
}

func (m *ObjectHelper) DelAttr(object Object, name string) {
	errors.ErrRuntime.Panic("not implemented: DelAttr")
}
