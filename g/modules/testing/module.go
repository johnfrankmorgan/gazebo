package testing

import (
	"github.com/johnfrankmorgan/gazebo/g"
)

type TestingModule struct {
	g.Base
	tests []*Test
	out   *g.Writer
	err   *g.Writer
}

func NewTestingModule() *TestingModule {
	object := &TestingModule{}
	object.SetSelf(object)
	return object
}

func (m *TestingModule) Value() interface{} {
	return m.Name()
}

func (m *TestingModule) Name() string {
	return "testing"
}

func (m *TestingModule) SetOutput(out, err *g.Writer) {
	m.out = out
	m.err = err
}

func (m *TestingModule) All() []*Test {
	return m.tests
}

// GAZEBO TESTING MODULE OBJECT METHODS

func (m *TestingModule) G_test(name g.Object, cb g.Object) *Test {
	test := NewTest(name.G_str().String(), cb)
	m.tests = append(m.tests, test)
	return test
}

func (m *TestingModule) G_run(tests ...*Test) {
	for _, test := range tests {
		if test.G_run().Bool() {
			m.out.Printf("PASS: %s\n", test.name)
		} else {
			m.err.Printf("FAIL: %s\n", test.name)
		}
	}
}
