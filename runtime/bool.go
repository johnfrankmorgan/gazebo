package runtime

type Bool bool

const (
	False = Bool(false)
	True  = Bool(true)
)

var BoolType = &Type{
	Name:   "Bool",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Hash:   func(self Object) uint64 { return self.(Bool).Hash() },
		Bool:   func(self Object) Bool { return self.(Bool).Bool() },
		String: func(self Object) String { return self.(Bool).String() },
	},
}

func (b Bool) Type() *Type {
	return BoolType
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

func (b Bool) String() String {
	if b {
		return "true"
	}

	return "false"
}
