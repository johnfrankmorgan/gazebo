package runtime

type Nil_ int

const Nil = Nil_(0)

func (Nil_) Type() *Type {
	return Types.Nil
}

func (Nil_) Bool() Bool {
	return False
}

func (Nil_) Repr() String {
	return "nil"
}
