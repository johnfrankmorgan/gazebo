package runtime

type Protocols struct {
	Hash   func(self Object) uint64
	Bool   func(self Object) Bool
	Repr   func(self Object) String
	String func(self Object) String

	Unary     UnaryProtocols
	Binary    BinaryProtocols
	Attribute AttributeProtocols
	Index     IndexProtocols
}

type UnaryProtocols struct {
	Positive func(self Object) Object
	Negative func(self Object) Object
}

type UnaryProtocol string

const (
	UnaryProtocolPositive UnaryProtocol = "+"
	UnaryProtocolNegative UnaryProtocol = "-"
)

type BinaryProtocols struct {
	Equal   func(self, other Object) Object
	Less    func(self, other Object) Object
	Greater func(self, other Object) Object

	Contains func(self, other Object) Object

	Add      func(self, other Object) Object
	Subtract func(self, other Object) Object
	Multiply func(self, other Object) Object
	Divide   func(self, other Object) Object
	Modulo   func(self, other Object) Object

	BitwiseAnd func(self, other Object) Object
	BitwiseOr  func(self, other Object) Object
	BitwiseXor func(self, other Object) Object

	ShiftLeft  func(self, other Object) Object
	ShiftRight func(self, other Object) Object

	Right RBinaryProtocols
}

type BinaryProtocol string

const (
	BinaryProtocolEqual   BinaryProtocol = "=="
	BinaryProtocolLess    BinaryProtocol = "<"
	BinaryProtocolGreater BinaryProtocol = ">"

	BinaryProtocolContains BinaryProtocol = "in"

	BinaryProtocolAdd      BinaryProtocol = "+"
	BinaryProtocolSubtract BinaryProtocol = "-"
	BinaryProtocolMultiply BinaryProtocol = "*"
	BinaryProtocolDivide   BinaryProtocol = "/"
	BinaryProtocolModulo   BinaryProtocol = "%"

	BinaryProtocolBitwiseAnd BinaryProtocol = "&"
	BinaryProtocolBitwiseOr  BinaryProtocol = "|"
	BinaryProtocolBitwiseXor BinaryProtocol = "^"

	BinaryProtocolShiftLeft  BinaryProtocol = "<<"
	BinaryProtocolShiftRight BinaryProtocol = ">>"
)

type RBinaryProtocols struct {
	Equal   func(self, other Object) Object
	Less    func(self, other Object) Object
	Greater func(self, other Object) Object

	Contains func(self, other Object) Object

	Add      func(self, other Object) Object
	Subtract func(self, other Object) Object
	Multiply func(self, other Object) Object
	Divide   func(self, other Object) Object
	Modulo   func(self, other Object) Object

	BitwiseAnd func(self, other Object) Object
	BitwiseOr  func(self, other Object) Object
	BitwiseXor func(self, other Object) Object

	ShiftLeft  func(self, other Object) Object
	ShiftRight func(self, other Object) Object
}

type AttributeProtocols struct {
	Get func(self Object, name String) Object
	Set func(self Object, name String, value Object)
}

type IndexProtocols struct {
	Get func(self, index Object) Object
	Set func(self, index, value Object)
}
