package vm

var _ Object = &NativeMethod{}

type NativeMethod struct {
	LazyAttributes
	value func(Object, Args) Object
}

func NewNativeMethod(f func(Object, Args) Object) *NativeMethod {
	return &NativeMethod{value: f}
}

func NewNativeMethodReturningBool(f func(Object, Args) *Bool) *NativeMethod {
	return NewNativeMethod(func(self Object, args Args) Object {
		return f(self, args)
	})
}

func NewNativeMethodReturningNumber(f func(Object, Args) *Number) *NativeMethod {
	return NewNativeMethod(func(self Object, args Args) Object {
		return f(self, args)
	})
}

func NewNativeMethodReturningString(f func(Object, Args) *String) *NativeMethod {
	return NewNativeMethod(func(self Object, args Args) Object {
		return f(self, args)
	})
}

func (m *NativeMethod) Type() Type {
	return Types.NativeMethod
}

func (m *NativeMethod) Value() interface{} {
	return m.value
}

func (m *NativeMethod) Func() func(Object, Args) Object {
	return m.value
}

func (m *NativeMethod) Bind(self Object) *BoundNativeMethod {
	return NewBoundNativeMethod(self, m)
}
