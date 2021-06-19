package vm

var _ Type = &TypeNativeMethod{}

type TypeNativeMethod struct {
	TypeBase
}

func (m *TypeNativeMethod) Name() *String {
	return NewString("NativeMethod")
}
