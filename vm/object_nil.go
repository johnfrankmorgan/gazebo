package vm

var _ Object = &Nil{}

type Nil struct {
	LazyAttributes
}

func NewNil() *Nil {
	return &Nil{}
}

func (m *Nil) Type() Type {
	return Types.Nil
}

func (m *Nil) Value() interface{} {
	return nil
}
