package vm

var _ Object = &Bool{}

type Bool struct {
	LazyAttributes
	value bool
}

func NewBool(value bool) Bool {
	return Bool{value: value}
}

func (m Bool) Type() Type {
	return Types.Bool
}

func (m Bool) Value() interface{} {
	return m.value
}

func (m Bool) Bool() bool {
	return m.value
}
