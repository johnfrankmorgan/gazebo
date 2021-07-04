package vm

var _ Type = &TypeBoundNativeMethod{}

type TypeBoundNativeMethod struct {
	TypeBase
}

func (m *TypeBoundNativeMethod) Name() String {
	return NewString("BoundNativeMethod")
}

func (m *TypeBoundNativeMethod) Call(self Object, args Args) Object {
	bound := self.(*BoundNativeMethod)
	return bound.Method().Func()(bound.Self(), args)
}
