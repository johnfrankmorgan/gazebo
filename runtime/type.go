package runtime

type Type struct {
	Name      string
	Parent    *Type
	Protocols TypeProtocols
	Ops       TypeOps
}

type TypeProtocols struct {
	Hash   func(self Object) uint64
	Bool   func(self Object) Bool
	String func(self Object) String
}

type TypeOps struct {
	Positive func(self Object) Object
	Negative func(self Object) Object

	Equal   func(self, other Object) Bool
	Less    func(self, other Object) Bool
	Greater func(self, other Object) Bool

	Contains func(self, other Object) Bool

	Add      func(self, other Object) Object
	Subtract func(self, other Object) Object
	Multiply func(self, other Object) Object
	Divide   func(self, other Object) Object
	Modulo   func(self, other Object) Object

	BitwiseAnd func(self, other Object) Object
	BitwiseOr  func(self, other Object) Object
	BitwiseXor func(self, other Object) Object

	LeftShift  func(self, other Object) Object
	RightShift func(self, other Object) Object
}

var TypeType = &Type{
	Name:   "Type",
	Parent: ObjectType,
}

func (t *Type) Type() *Type {
	return TypeType
}
