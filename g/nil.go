package g

var _ Object = &Nil{}

type Nil struct {
	Base
}

func NewNil() *Nil {
	object := &Nil{}
	object.self = object
	return object
}

func (m *Nil) Value() interface{} {
	return nil
}

// GAZEBO NIL OBJECT METHODS

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
