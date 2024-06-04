package runtime

type Bool bool

const (
	False = Bool(false)
	True  = Bool(true)
)

func (b Bool) Type() *Type {
	return Types.Bool
}

func (b Bool) Hash() uint64 {
	if b {
		return 1
	}

	return 0
}

func (b Bool) Bool() Bool {
	return b
}

func (b Bool) Repr() String {
	if b {
		return "true"
	}

	return "false"
}
