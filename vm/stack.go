package vm

import (
	"github.com/johnfrankmorgan/gazebo/assert"
	"github.com/johnfrankmorgan/gazebo/debug"
	"github.com/johnfrankmorgan/gazebo/g"
)

type stack struct {
	values []g.Object
}

func (m *stack) dump() {
	debug.Printf("STACK\n")

	for off, value := range m.values {
		debug.Printf("%4d: %v\n", m.size()-off, value.Value())
	}

	debug.Printf("\n")
}

func (m *stack) top() g.Object {
	if size := m.size(); size > 0 {
		return m.values[size-1]
	}

	return nil
}

func (m *stack) push(value g.Object) {
	m.values = append(m.values, value)
}

func (m *stack) peek() g.Object {
	if size := m.size(); size > 0 {
		value := m.values[size-1]
		return value
	}

	assert.Unreached("stack empty")
	return nil
}

func (m *stack) pop() g.Object {
	if size := m.size(); size > 0 {
		value := m.values[size-1]
		m.values = m.values[:size-1]
		return value
	}

	assert.Unreached("stack empty")
	return nil
}

func (m *stack) size() int {
	return len(m.values)
}
