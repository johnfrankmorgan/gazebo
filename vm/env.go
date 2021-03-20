package vm

import (
	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g"
)

type env struct {
	parent *env
	values *g.Map
}

func (m *env) init() {
	if m.values == nil {
		m.values = g.NewMap()
	}
}

func (m *env) resolve(name string) *env {
	m.init()

	if m.values.HasAttr(name) {
		return m
	}

	if m.parent != nil {
		return m.parent.resolve(name)
	}

	return nil
}

func (m *env) lookup(name string) g.Object {
	m.init()

	if env := m.resolve(name); env != nil {
		return env.values.GetAttr(name)
	}

	errors.ErrRuntime.Panic("undefined name: %s", name)
	return nil
}

func (m *env) defined(name string) bool {
	m.init()

	return m.resolve(name) != nil
}

func (m *env) define(name string, value g.Object) {
	m.init()

	m.values.SetAttr(name, value)
}

func (m *env) assign(name string, value g.Object) {
	m.init()

	if env := m.resolve(name); env != nil {
		env.values.SetAttr(name, value)
		return
	}

	errors.ErrRuntime.Panic("undefined name: %s", name)
	return
}

func (m *env) remove(name string) {
	m.init()

	if env := m.resolve(name); env != nil {
		env.values.DelAttr(name)
	}
}
