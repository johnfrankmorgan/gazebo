package vm

import "github.com/johnfrankmorgan/gazebo/runtime"

type Stack []runtime.Object

func (s Stack) Size() int {
	return len(s)
}

func (s *Stack) Push(object runtime.Object) {
	*s = append(*s, object)
}

func (s *Stack) Pop() runtime.Object {
	object := (*s)[s.Size()-1]
	*s = (*s)[:len(*s)-1]

	return object
}
