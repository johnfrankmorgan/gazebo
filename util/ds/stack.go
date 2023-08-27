package ds

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Size() int {
	return len(s.values)
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() T {
	value := s.Peek()
	s.values = s.values[:s.Size()-1]
	return value
}

func (s *Stack[T]) Peek() T {
	return s.values[s.Size()-1]
}
