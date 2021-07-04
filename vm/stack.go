package vm

type Stack struct {
	values []Object
}

func NewStack() *Stack {
	return &Stack{}
}

func (m *Stack) Size() int {
	return len(m.values)
}

func (m *Stack) Push(value Object) {
	m.values = append(m.values, value)
}

func (m *Stack) Peek() Object {
	return m.values[m.Size()-1]
}

func (m *Stack) Pop() Object {
	defer func() { m.values = m.values[:m.Size()-1] }()
	return m.Peek()
}
