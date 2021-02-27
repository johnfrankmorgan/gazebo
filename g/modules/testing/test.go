package testing

import (
	"github.com/johnfrankmorgan/gazebo/g"
)

var _ g.Object = &Test{}

type Test struct {
	g.Base
	name   string
	cb     g.Object
	failed bool
}

func NewTest(name string, cb g.Object) *Test {
	object := &Test{
		name: name,
		cb:   cb,
	}
	object.SetSelf(object)
	return object
}

func (m *Test) Value() interface{} {
	return m.name
}

func (m *Test) Name() string {
	return m.name
}

func (m *Test) Failed() bool {
	return m.failed
}

// GAZEBO TEST OBJECT METHODS

func (m *Test) G_run() *g.Bool {
	success := m.cb.G_invoke(g.NewArgs(nil)).G_bool()
	m.failed = success.G_not().Bool()
	return success
}
