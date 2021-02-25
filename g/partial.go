package g

import (
	"reflect"

	"github.com/johnfrankmorgan/gazebo/errors"
)

var _ Protocols = &Partial{}

type Partial struct {
	self Object
}

func (m *Partial) G_str() *String {
	return NewStringf("<%T>(%v)", m.self, m.self)
}

func (m *Partial) G_num() *Number {
	return NewNumber(0)
}

func (m *Partial) G_bool() *Bool {
	return NewBool(true)
}

func (m *Partial) G_not() *Bool {
	return NewBool(!m.G_bool().Bool())
}

func (m *Partial) G_eq(other Object) *Bool {
	return NewBool(reflect.DeepEqual(m.self.Value(), other.Value()))
}

func (m *Partial) G_neq(other Object) *Bool {
	return NewBool(!m.G_eq(other).Bool())
}

func (m *Partial) G_gt(other Object) *Bool {
	errors.ErrRuntime.Panic("gt not implemented for %T", m.self)
	return nil
}

func (m *Partial) G_gte(other Object) *Bool {
	errors.ErrRuntime.Panic("gte not implemented for %T", m.self)
	return nil
}

func (m *Partial) G_lt(other Object) *Bool {
	errors.ErrRuntime.Panic("lt not implemented for %T", m.self)
	return nil
}

func (m *Partial) G_lte(other Object) *Bool {
	errors.ErrRuntime.Panic("lte not implemented for %T", m.self)
	return nil
}
