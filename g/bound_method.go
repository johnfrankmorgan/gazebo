package g

var _ Object = &BoundMethod{}

type BoundMethod struct {
	Base
	self   Object
	name   string
	method Method
}

func NewBoundMethod(self Object, name string, method Method) *BoundMethod {
	object := &BoundMethod{
		self:   self,
		name:   name,
		method: method,
	}

	object.SetType(TypeBoundMethod)
	object.SetSelf(object)

	return object
}

func (m *BoundMethod) Value() interface{} {
	return m.method
}

func (m *BoundMethod) ToString() *String {
	return NewStringf(
		"%s(%s.%s@%v self@%p = %q)",
		m.Type().Name(),
		m.self.Type().Name(),
		m.name,
		m.method,
		m.self,
		m.self.ToString().Limit(10, "...").String(),
	)
}

func (m *BoundMethod) Call(args *Args) Object {
	return m.method(m.self, args)
}
