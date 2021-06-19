package vm

var Types struct {
	Type              *TypeType
	Nil               *TypeNil
	Bool              *TypeBool
	Number            *TypeNumber
	String            *TypeString
	Map               *TypeMap
	NativeMethod      *TypeNativeMethod
	BoundNativeMethod *TypeBoundNativeMethod
}

func init() {
	Types.Type = &TypeType{}
	Types.Type.SetSelf(Types.Type)

	Types.Nil = &TypeNil{}
	Types.Nil.SetSelf(Types.Nil)

	Types.Bool = &TypeBool{}
	Types.Bool.SetSelf(Types.Bool)

	Types.Number = &TypeNumber{}
	Types.Number.SetSelf(Types.Number)

	Types.String = &TypeString{}
	Types.String.SetSelf(Types.String)

	Types.Map = &TypeMap{}
	Types.Map.SetSelf(Types.Map)

	Types.NativeMethod = &TypeNativeMethod{}
	Types.NativeMethod.SetSelf(Types.NativeMethod)

	Types.BoundNativeMethod = &TypeBoundNativeMethod{}
	Types.BoundNativeMethod.SetSelf(Types.BoundNativeMethod)
}

type Type interface {
	Object

	Name() *String
	Methods() map[string]Object

	ToBool(Object, Args) *Bool
	IsNil(Object, Args) *Bool
	ToString(Object, Args) *String
	ToNumber(Object, Args) *Number

	Eq(Object, Args) *Bool
	NEq(Object, Args) *Bool
	Gt(Object, Args) *Bool
	GtE(Object, Args) *Bool
	Lt(Object, Args) *Bool
	LtE(Object, Args) *Bool

	Add(Object, Args) Object
	Sub(Object, Args) Object
	Mul(Object, Args) Object
	Div(Object, Args) Object

	HasAttr(Object, Args) *Bool
	GetAttr(Object, Args) Object
	SetAttr(Object, Args) Object
	DelAttr(Object, Args) Object

	Hash(Object, Args) *Number
	Call(Object, Args) Object
}
