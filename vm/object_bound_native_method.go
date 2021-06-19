package vm

var _ Object = &BoundNativeMethod{}

type BoundNativeMethod struct {
	LazyAttributes
	self   Object
	method *NativeMethod
}

func NewBoundNativeMethod(self Object, method *NativeMethod) *BoundNativeMethod {
	return &BoundNativeMethod{
		self:   self,
		method: method,
	}
}

func (m *BoundNativeMethod) Type() Type {
	return Types.BoundNativeMethod
}

func (m *BoundNativeMethod) Value() interface{} {
	return struct {
		Self   Object
		Method *NativeMethod
	}{m.Self(), m.Method()}
}

func (m *BoundNativeMethod) Self() Object {
	return m.self
}

func (m *BoundNativeMethod) Method() *NativeMethod {
	return m.method
}
