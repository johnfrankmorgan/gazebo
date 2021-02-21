package modules

import (
	"github.com/johnfrankmorgan/gazebo/g"
)

// All returns all available modules
func All() map[string]*Module {
	return map[string]*Module{
		"http": HTTP,
		"str":  Str,
		"time": Time,
	}
}

var TypeModule *g.Type

func init() {
	TypeModule = &g.Type{
		Name:    "Module",
		Parent:  g.TypeBase,
		Methods: g.Methods{},
	}
}

type ObjectModule struct {
	g.PartialObject
	value      *Module
	attributes g.Attributes
}

func (m *ObjectModule) Value() interface{} {
	return m.value
}

func (m *ObjectModule) Call(method string, args g.Args) g.Object {
	return m.CallMethod(m, method, args)
}

func (m *ObjectModule) Attributes() *g.Attributes {
	return &m.attributes
}

type Module struct {
	Name   string
	Init   func(*Module)
	Values map[string]g.Object
}

func (m *Module) Load() g.Object {
	m.Init(m)

	object := &ObjectModule{
		value: m,
	}

	object.SetType(TypeModule)

	for name, value := range m.Values {
		object.attributes.Set(name, value)
	}

	return object
}
