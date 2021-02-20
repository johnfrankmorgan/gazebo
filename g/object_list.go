package g

// ObjectList is the underlying type of lists in gazebo
type ObjectList struct {
	PartialObject
	value []Object
}

// NewObjectList creates a new list object
func NewObjectList(value []Object) *ObjectList {
	return &ObjectList{
		PartialObject: PartialObject{typ: TypeList},
		value:         value,
	}
}

// Value satisfies the Object interface
func (m *ObjectList) Value() interface{} {
	return m.value
}

// Call satisfies the Object interface
func (m *ObjectList) Call(method string, args Args) Object {
	return m.call(m, method, args)
}

// Slice returns the list's underlying slice
func (m *ObjectList) Slice() []Object {
	return m.value
}

// Append appends an Object to the List
func (m *ObjectList) Append(objects ...Object) {
	m.value = append(m.value, objects...)
}

// Set sets the value at an index
func (m *ObjectList) Set(index int, value Object) {
	m.value[index] = value
}

// Index returns the Object at the provided index
func (m *ObjectList) Index(index int) Object {
	return m.value[index]
}

// Len returns the length of the list
func (m *ObjectList) Len() int {
	return len(m.value)
}
