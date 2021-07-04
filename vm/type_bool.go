package vm

var _ Type = &TypeBool{}

type TypeBool struct {
	TypeBase
}

func (m *TypeBool) Name() String {
	return NewString("Bool")
}

func (m *TypeBool) ToString(self Object, _ Args) String {
	if self.(Bool).Bool() {
		return NewString("true")
	}

	return NewString("false")
}

func (m *TypeBool) ToNumber(self Object, _ Args) Number {
	if self.(Bool).Bool() {
		return NewNumber(1.0)
	}

	return NewNumber(0.0)
}

func (m *TypeBool) ToBool(self Object, _ Args) Bool {
	return self.(Bool)
}

func (m *TypeBool) Eq(self Object, args Args) Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToBool(args[0], nil).Bool()
	return NewBool(self.(*Bool).Bool() == other)
}

func (m *TypeBool) NEq(self Object, args Args) Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToBool(args[0], nil).Bool()
	return NewBool(self.(*Bool).Bool() != other)
}
