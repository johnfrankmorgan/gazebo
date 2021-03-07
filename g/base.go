package g

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/maphash"
	"reflect"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/g/protocols"
)

var (
	_ Protocols = &Base{}
	_ Attrs     = &Base{}
)

type Base struct {
	self  Object
	attrs Attributes
}

func (m *Base) unimplemented(name string) {
	errors.ErrRuntime.Panic("%q not implemented for %T", name, m.self)
}

func (m *Base) Method(object Object, name string) *BoundMethod {
	rvalue := reflect.ValueOf(object)
	length := len(name)

	if name[length-1] == '?' {
		name = name[:length-1]
	}

	method := rvalue.MethodByName("G_" + name)

	if method.IsValid() {
		return NewBoundMethod(
			fmt.Sprintf("%T", object),
			name,
			method,
		)
	}

	return nil
}

func (m *Base) SetSelf(self Object) {
	m.self = self
}

// OBJECT METHODS

func (m *Base) CallMethod(name string, args *Args) Object {
	if method := m.Method(m.self, name); method != nil {
		return method.G_invoke(args)
	}

	m.unimplemented(name)
	return nil
}

var _hash maphash.Hash

func (m *Base) Hash() uint64 {
	var buff bytes.Buffer

	err := gob.NewEncoder(&buff).Encode(m.self.Value())
	errors.ErrRuntime.ExpectNilError(err)

	defer _hash.Reset()

	_hash.Write(buff.Bytes())

	return _hash.Sum64()
}

// ATTRIBUTE METHODS

func (m *Base) HasAttr(name string) bool {
	return m.attrs.Has(name)
}

func (m *Base) GetAttr(name string) Object {
	if m.HasAttr(name) {
		return m.attrs.Get(name)
	}

	if method := m.Method(m.self, name); method != nil {
		return method
	}

	errors.ErrRuntime.Panic("attribute %q undefined for %T", name, m.self)
	return nil
}

func (m *Base) SetAttr(name string, value Object) {
	m.attrs.Set(name, value)
}

func (m *Base) DelAttr(name string) {
	m.attrs.Delete(name)
}

// PROTOCOLS

func (m *Base) G_repr() *String {
	return NewStringf("<%T>(%v)", m.self, m.self)
}

func (m *Base) G_str() *String {
	return m.self.G_repr()
}

func (m *Base) G_num() *Number {
	m.unimplemented(protocols.Number)
	return nil
}

func (m *Base) G_bool() *Bool {
	return NewBool(true)
}

func (m *Base) G_not() *Bool {
	return NewBool(!m.self.G_bool().Bool())
}

func (m *Base) G_len() *Number {
	m.unimplemented(protocols.Len)
	return nil
}

func (m *Base) G_inverse() Object {
	m.unimplemented(protocols.Inverse)
	return nil
}

func (m *Base) G_and(other Object) *Bool {
	return NewBool(m.self.G_bool().Bool() && other.G_bool().Bool())
}

func (m *Base) G_or(other Object) Object {
	if m.self.G_bool().Bool() {
		return m.self
	}

	return other
}

func (m *Base) G_contains(_ Object) *Bool {
	m.unimplemented(protocols.Contains)
	return nil
}

func (m *Base) G_add(_ Object) Object {
	m.unimplemented(protocols.Add)
	return nil
}

func (m *Base) G_sub(_ Object) Object {
	m.unimplemented(protocols.Sub)
	return nil
}

func (m *Base) G_mul(_ Object) Object {
	m.unimplemented(protocols.Mul)
	return nil
}

func (m *Base) G_div(_ Object) Object {
	m.unimplemented(protocols.Div)
	return nil
}

func (m *Base) G_eq(other Object) *Bool {
	return NewBool(reflect.DeepEqual(m.self.Value(), other.Value()))
}

func (m *Base) G_neq(other Object) *Bool {
	return NewBool(!m.self.G_eq(other).Bool())
}

func (m *Base) G_gt(_ Object) *Bool {
	m.unimplemented(protocols.GreaterThan)
	return nil
}

func (m *Base) G_gte(_ Object) *Bool {
	m.unimplemented(protocols.GreaterThanEqual)
	return nil
}

func (m *Base) G_lt(_ Object) *Bool {
	m.unimplemented(protocols.LessThan)
	return nil
}

func (m *Base) G_lte(_ Object) *Bool {
	m.unimplemented(protocols.LessThanEqual)
	return nil
}

func (m *Base) G_hasattr(name *String) *Bool {
	return NewBool(m.self.HasAttr(name.String()))
}

func (m *Base) G_getattr(name *String) Object {
	return m.self.GetAttr(name.String())
}

func (m *Base) G_setattr(name *String, value Object) Object {
	m.self.SetAttr(name.String(), value)
	return NewNil()
}

func (m *Base) G_delattr(name *String) Object {
	m.self.DelAttr(name.String())
	return NewNil()
}

func (m *Base) G_invoke(_ *Args) Object {
	m.unimplemented(protocols.Invoke)
	return nil
}
