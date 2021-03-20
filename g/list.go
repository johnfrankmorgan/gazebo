package g

var _ Object = &List{}

type List struct {
	Base
	value []Object
}

func NewList(value []Object) *List {
	object := &List{value: value}

	object.SetType(TypeList)
	object.SetSelf(object)

	return object
}

func NewListSized(size int) *List {
	return NewList(make([]Object, size, 0))
}

func (m *List) Value() interface{} {
	return m.value
}

func (m *List) Slice() []Object {
	return m.value
}

func (m *List) Len() int {
	return len(m.value)
}

func (m *List) Get(offset int) Object {
	return m.value[offset]
}

func (m *List) Set(offset int, value Object) {
	m.value[offset] = value
}

func (m *List) Prepend(value Object) {
	m.value = append([]Object{value}, m.value...)
}

func (m *List) Append(value Object) {
	m.value = append(m.value, value)
}
