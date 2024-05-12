package runtime

type Float float64

var FloatType = &Type{
	Name:   "Float",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Hash:   func(self Object) uint64 { return self.(Float).Hash() },
		Bool:   func(self Object) Bool { return self.(Float).Bool() },
		String: func(self Object) String { return self.(Float).String() },
	},
}

func (f Float) Type() *Type {
	return FloatType
}

func (f Float) Hash() uint64 {
	return hash(f)
}

func (f Float) Bool() Bool {
	return f != 0
}

func (f Float) String() String {
	return Stringf("%f", f)
}
