package inspect

import (
	"reflect"
	"strings"

	"github.com/johnfrankmorgan/gazebo/g"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
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
			list.Append(object.GetAttr(name[2:]))
		}
	}

	return list
}

func (m *InspectModule) G_protocol(object g.Object) *g.Bool {
	var name string

	if method, ok := object.(*g.BoundMethod); ok {
		name = method.G_name().String()
	} else {
		name = object.G_str().String()
	}

	for _, p := range protocols.All() {
		if p == name {
			return g.NewBool(true)
		}
	}

	return g.NewBool(false)
}

func (m *InspectModule) G_protocols() *g.List {
	protocols := protocols.All()

	list := g.NewListSized(len(protocols))

	for i, p := range protocols {
		list.Set(i, g.NewString(p))
	}

	return list
}
