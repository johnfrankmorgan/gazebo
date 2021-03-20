package g

var _ Object = &BoundMethod{}

type BoundMethod struct {
	Base
	self   Object
	method Method
}

func NewBoundMethod(self Object, method Method) *BoundMethod {
	object := &BoundMethod{
		self:   self,
		method: method,
	}

	object.SetType(TypeBoundMethod)
	object.SetSelf(object)

	return object
}

func (m *BoundMethod) Value() interface{} {
	return m.method
}

func (m *BoundMethod) Call(args *Args) Object {
	return m.method(m.self, args)
}
