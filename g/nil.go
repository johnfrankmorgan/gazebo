package g

var _ Object = &Nil{}

type Nil struct {
	Partial
	h ObjectHelper
}

func NewNil() *Nil {
	object := &Nil{}
	object.self = object
	return object
}

func (m *Nil) Value() interface{} {
	return nil
}

func (m *Nil) CallMethod(name string, args *Args) Object {
	return m.h.CallMethod(m, name, args)
}

func (m *Nil) HasAttr(name string) bool {
	return m.h.HasAttr(m, name)
}

func (m *Nil) GetAttr(name string) Object {
	return m.h.GetAttr(m, name)
}

func (m *Nil) SetAttr(name string, value Object) {
	m.h.SetAttr(m, name, value)
}

func (m *Nil) DelAttr(name string) {
	m.h.DelAttr(m, name)
}

// GAZEBO STRING OBJECT METHODS

func (m *Nil) G_str() *String {
	return NewString("nil")
}

func (m *Nil) G_num() *Number {
	return NewNumber(0)
}

func (m *Nil) G_bool() *Bool {
	return NewBool(false)
}

func (m *Nil) G_not() *Bool {
	return NewBool(true)
}
