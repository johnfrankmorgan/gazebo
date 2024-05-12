package runtime

type Nil_ int

const Nil = Nil_(0)

var NilType = &Type{
	Name:   "Nil",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Bool:   func(self Object) Bool { return self.(Nil_).Bool() },
		String: func(self Object) String { return self.(Nil_).String() },
	},
}

func (n Nil_) Type() *Type {
	return NilType
}

func (n Nil_) Bool() Bool {
	return False
}

func (n Nil_) String() String {
	return "nil"
}
