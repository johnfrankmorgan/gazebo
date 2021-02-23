package g

import (
	"reflect"

	"github.com/johnfrankmorgan/gazebo/assert"
)

type Object interface {
	Attrs
	Value() interface{}
}

type ObjectHelper struct {
	Attrs map[string]Object
}

func (m *ObjectHelper) Method(object Object, name string) Object {
	method := reflect.ValueOf(object).MethodByName("G_" + name)

	if method.IsValid() {
		return NewBoundMethod(method)
	}

	return nil
}

func (m *ObjectHelper) HasAttr(object Object, name string) bool {
	assert.Unreached()
	return false
}

func (m *ObjectHelper) GetAttr(object Object, name string) Object {
	if method := m.Method(object, name); method != nil {
		return method
	}

	assert.Unreached()
	return nil
}

func (m *ObjectHelper) SetAttr(object Object, name string, value Object) {
	assert.Unreached()
}

func (m *ObjectHelper) DelAttr(object Object, name string) {
	assert.Unreached()
}
