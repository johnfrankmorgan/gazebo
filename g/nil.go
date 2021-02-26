package g

var _ Object = &Nil{}

type Nil struct {
	Base
}

func NewNil() *Nil {
	object := &Nil{}
	object.SetSelf(object)
	return object
}

func (m *Nil) Value() interface{} {
	return nil
}

// GAZEBO NIL OBJECT PROTOCOLS

func (m *Nil) G_repr() *String {
	return m.G_str()
}

func (m *Nil) G_str() *String {
	return NewString("nil")
}

func (m *Nil) G_num() *Number {
	return NewNumber(0)
}

func (m *Nil) G_bool() *Bool {
	return NewBool(false)
}
