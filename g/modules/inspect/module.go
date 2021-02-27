package inspect

import (
	"reflect"
	"strings"

	"github.com/johnfrankmorgan/gazebo/g"
)

type InspectModule struct {
	g.Base
}

func NewInspectModule() *InspectModule {
	object := &InspectModule{}
	object.SetSelf(object)
	return object
}

func (m *InspectModule) Name() string {
	return "inspect"
}

func (m *InspectModule) Value() interface{} {
	return m.Name()
}

// GAZEBO OS MODULE OBJECT METHODS

func (m *InspectModule) G_type(object g.Object) *g.String {
	return g.NewStringf("%T", object)
}

func (m *InspectModule) G_methods(object g.Object) *g.List {
	typ := reflect.TypeOf(object)
	list := g.NewList(nil)

	for i := 0; i < typ.NumMethod(); i++ {
		name := typ.Method(i).Name

		if strings.HasPrefix(name, "G_") {
			list.Append(object.GetAttr(strings.Replace(name, "G_", "", 1)))
		}
	}

	return list
}
