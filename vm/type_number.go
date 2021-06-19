package vm

var _ Type = &TypeNumber{}

type TypeNumber struct {
	TypeBase
}

func (m *TypeNumber) Name() *String {
	return NewString("Number")
}

func (m *TypeNumber) ToBool(self Object, _ Args) *Bool {
	return NewBool(self.(*Number).Float() != 0.0)
}

func (m *TypeNumber) ToNumber(self Object, _ Args) *Number {
	return self.(*Number)
}

func (m *TypeNumber) ToString(self Object, _ Args) *String {
	return NewStringf("%f", self.(*Number).Float())
}

func (m *TypeNumber) Eq(self Object, args Args) *Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewBool(self.(*Number).Float() == other)
}

func (m *TypeNumber) NEq(self Object, args Args) *Bool {

	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewBool(self.(*Number).Float() != other)
}

func (m *TypeNumber) Gt(self Object, args Args) *Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewBool(self.(*Number).Float() > other)

}

func (m *TypeNumber) GtE(self Object, args Args) *Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewBool(self.(*Number).Float() >= other)
}

func (m *TypeNumber) Lt(self Object, args Args) *Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewBool(self.(*Number).Float() < other)
}

func (m *TypeNumber) LtE(self Object, args Args) *Bool {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewBool(self.(*Number).Float() <= other)
}

func (m *TypeNumber) Add(self Object, args Args) Object {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewNumber(self.(*Number).Float() + other)
}

func (m *TypeNumber) Sub(self Object, args Args) Object {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewNumber(self.(*Number).Float() - other)
}

func (m *TypeNumber) Mul(self Object, args Args) Object {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewNumber(self.(*Number).Float() * other)
}

func (m *TypeNumber) Div(self Object, args Args) Object {
	args.ExpectsExactly(1)

	other := args[0].Type().ToNumber(args[0], nil).Float()
	return NewNumber(self.(*Number).Float() / other)
}
