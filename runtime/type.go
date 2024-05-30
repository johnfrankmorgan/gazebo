package runtime

type Type struct {
	Name       string
	Parent     *Type
	Protocols  TypeProtocols
	Ops        TypeOps
	Attributes TypeAttributes
}

type TypeProtocols struct {
	Hash   func(self Object) uint64
	Bool   func(self Object) Bool
	String func(self Object) String
}

type TypeOps struct {
	Positive func(self Object) Object
	Negative func(self Object) Object

	GetAttribute func(self Object, name String) Object
	SetAttribute func(self Object, name String, value Object)

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

	GetIndex func(self, index Object) Object
	SetIndex func(self, index, value Object)
}

type TypeAttributes map[String]Attribute

type Attribute struct {
	Get func(self Object) Object
	Set func(self, value Object)
}

var TypeType = &Type{
	Name:   "Type",
	Parent: ObjectType,
	Attributes: TypeAttributes{
		"name": Attribute{
			Get: func(self Object) Object { return String(self.(*Type).Name) },
			Set: func(self, value Object) { self.(*Type).Name = string(value.(String)) },
		},

		"parent": Attribute{
			Get: func(self Object) Object {
				if parent := self.(*Type).Parent; parent != nil {
					return parent
				}

				return Nil
			},
		},
	},
}

func (t *Type) Type() *Type {
	return TypeType
}
