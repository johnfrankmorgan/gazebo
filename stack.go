package gazebo

type stack[T any] struct {
	values []T
}

func (s *stack[T]) size() int {
	return len(s.values)
}

func (s *stack[T]) top() *T {
	return &s.values[s.size()-1]
}

func (s *stack[T]) peek() T {
	return s.values[s.size()-1]
}

func (s *stack[T]) push(value T) {
	s.values = append(s.values, value)
}

func (s *stack[T]) pop() T {
	value := s.peek()
	s.values = s.values[:s.size()-1]
	return value
}
